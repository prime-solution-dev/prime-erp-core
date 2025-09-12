package quotationService

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"prime-erp-core/internal/db"
	"prime-erp-core/internal/models"
	approvalService "prime-erp-core/internal/services/approval-service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateQuotationRequest struct {
	IsVerifyPrice bool                `json:"is_verify_price"` // true = verify, if not verified can't create
	Quotations    []QuotationDocument `json:"quotations"`
}

type QuotationDocument struct {
	models.Quotation
	Items []models.QuotationItem `json:"items"`
}

type CreateQuotationResponse struct {
	IsPass        bool   `json:"is_pass"`
	QuotationCode string `json:"quotation_code"`
}

func CreateQuotation(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	req := CreateQuotationRequest{}
	res := []CreateQuotationResponse{}

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	sqlx, err := db.ConnectSqlx(`prime_erp`)
	if err != nil {
		return nil, err
	}
	defer sqlx.Close()

	gormx, err := db.ConnectGORM(`prime_erp`)
	if err != nil {
		return nil, err
	}
	defer db.CloseGORM(gormx)

	user := `system` // TODO: get from ctx
	now := time.Now()
	nowTruc := now.Truncate(24 * time.Hour)

	createQuotations := []models.Quotation{}
	createQuotationItems := []models.QuotationItem{}
	verifyReqMap := map[string]approvalService.VerifyApproveRequest{}

	for _, quotationReq := range req.Quotations {
		tempQuotation := quotationReq.Quotation
		tempQuotation.ID = uuid.New()

		if quotationReq.EffectiveDatePrice == nil {
			return nil, fmt.Errorf("effective date is required for quotation %s", quotationReq.QuotationCode)
		}

		if tempQuotation.QuotationCode == "" {
			tempQuotation.QuotationCode = uuid.New().String()
		}

		tempQuotation.CreateDate = &now
		tempQuotation.CreateBy = user
		tempQuotation.UpdateDate = &now
		tempQuotation.UpdateBy = user

		createQuotations = append(createQuotations, tempQuotation)

		//Approval
		verifyReqKey := fmt.Sprintf(`%s|%s`, quotationReq.CompanyCode, quotationReq.SiteCode)
		verifyReq, existVerifyReq := verifyReqMap[verifyReqKey]
		if !existVerifyReq {
			newVerifyReq := approvalService.VerifyApproveRequest{
				IsVerifyPrice: true,
				CompanyCode:   quotationReq.CompanyCode,
				SiteCode:      quotationReq.SiteCode,
				StorageType:   []string{`NORMAL`},
				SaleDate:      nowTruc,
			}

			verifyReq = newVerifyReq
		}

		newApprDoc := approvalService.VerifyApproveDocument{
			DocRef:             quotationReq.QuotationCode,
			CustomerCode:       quotationReq.CustomerCode,
			EffectiveDatePrice: *quotationReq.EffectiveDatePrice,
			Items:              []approvalService.VerifyApproveItem{},
		}

		for _, item := range quotationReq.Items {
			item.ID = uuid.New()
			item.QuotationID = tempQuotation.ID

			if item.QuotationItem == "" {
				item.QuotationItem = uuid.New().String()
			}

			item.CreateDate = &now
			item.CreateBy = user
			item.UpdateDate = &now
			item.UpdateBy = user

			createQuotationItems = append(createQuotationItems, item)

			//Approval
			newApprItem := approvalService.VerifyApproveItem{
				ItemRef:       item.QuotationItem,
				ProductCode:   item.ProductCode,
				Qty:           item.Qty,
				Unit:          item.Unit,
				TotalWeight:   item.TotalWeight,
				PriceUnit:     item.PriceUnit,
				PriceListUnit: item.PriceListUnit,
				TotalAmount:   item.TotalAmount,
				SaleUnit:      item.SaleUnit,
				SaleUnitType:  item.SaleUnitType,
			}

			newApprDoc.Items = append(newApprDoc.Items, newApprItem)
		}

		//Approval
		verifyReq.Documents = append(verifyReq.Documents, newApprDoc)
		verifyReqMap[verifyReqKey] = verifyReq
	}

	//Verification
	if req.IsVerifyPrice {
		for _, verifyReq := range verifyReqMap {
			verifyRes, err := approvalService.VerifyApproveLogic(gormx, sqlx, verifyReq)
			if err != nil {
				return nil, err
			}

			for _, doc := range verifyRes.Documents {
				res = append(res, CreateQuotationResponse{
					IsPass:        doc.IsPassPrice,
					QuotationCode: doc.DocRef,
				})
			}
		}
	}

	// check duplicate quotation codes
	var existCount int64
	codes := make([]string, 0, len(createQuotations))
	for _, q := range createQuotations {
		codes = append(codes, q.QuotationCode)
	}

	tx := gormx.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if len(codes) > 0 {
		if err := tx.Model(&models.Quotation{}).
			Where("quotation_code IN ?", codes).
			Count(&existCount).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		if existCount > 0 {
			tx.Rollback()
			return nil, errors.New("duplicate quotation code detected")
		}
	}

	// Insert quotations
	if len(createQuotations) > 0 {
		if err := tx.Create(&createQuotations).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Insert items
	if len(createQuotationItems) > 0 {
		if err := tx.Create(&createQuotationItems).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return res, nil
}
