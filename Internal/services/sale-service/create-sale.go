package saleService

import (
	"encoding/json"
	"errors"
	"fmt"
	"prime-erp-core/Internal/models"
	"prime-erp-core/internal/db"
	approvalService "prime-erp-core/internal/services/approval-service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateSaleRequest struct {
	IsVerifyPrice      bool
	IsVerifyCredit     bool
	IsVerifyExpiryDate bool
	IsVerifyInventory  bool
	Sales              []SaleDocument
}

type SaleDocument struct {
	models.Sale
	Items []models.SaleItem
}

type CreateSaleResponse struct {
	IsPass           bool   `json:"is_pass"`
	IsPassPrice      bool   `json:"is_pass_price"`
	IsPassCredit     bool   `json:"is_pass_credit"`
	IsPassInventory  bool   `json:"is_pass_inventory"`
	IsPassExpiryDate bool   `json:"is_pass_expiry_date"`
	SaleCode         string `json:"sale_code"`
}

func CreateSale(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	req := CreateSaleRequest{}
	res := []CreateSaleResponse{}

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

	createSales := []models.Sale{}
	createSaleItems := []models.SaleItem{}
	verifyReqMap := map[string]approvalService.VerifyApproveRequest{}

	for _, saleReq := range req.Sales {
		tempSale := saleReq.Sale
		tempSale.ID = uuid.New()

		if tempSale.SaleCode == "" {
			tempSale.SaleCode = uuid.New().String()
		}

		tempSale.CreateDate = &now
		tempSale.CreateBy = user
		tempSale.UpdateDate = &now
		tempSale.UpdateBy = user

		createSales = append(createSales, tempSale)

		//Approval
		verifyReqKey := fmt.Sprintf(`%s|%s`, saleReq.CompanyCode, saleReq.SiteCode)
		verifyReq, existVerifyReq := verifyReqMap[verifyReqKey]
		if !existVerifyReq {
			newVerifyReq := approvalService.VerifyApproveRequest{
				IsVerifyPrice:       req.IsVerifyPrice,
				IsVerifyCredit:      req.IsVerifyCredit,
				IsVerifyExpiryPrice: req.IsVerifyExpiryDate,
				IsVerifyInventory:   req.IsVerifyInventory,
				CompanyCode:         saleReq.CompanyCode,
				SiteCode:            saleReq.SiteCode,
				StorageType:         []string{`NORMAL`},
				SaleDate:            nowTruc,
			}
			verifyReq = newVerifyReq
		}

		newApprDoc := approvalService.VerifyApproveDocument{
			DocRef:       tempSale.SaleCode,
			CustomerCode: saleReq.CustomerCode,
			Items:        []approvalService.VerifyApproveItem{},
		}

		for _, item := range saleReq.Items {
			item.ID = uuid.New()
			item.SaleID = tempSale.ID
			item.CreateDate = &now
			item.CreateBy = user
			item.UpdateDate = &now
			item.UpdateBy = user

			createSaleItems = append(createSaleItems, item)

			//Approval
			newApprItem := approvalService.VerifyApproveItem{
				ItemRef:       item.SaleItem,
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

	// Verification
	for _, verifyReq := range verifyReqMap {
		verifyRes, err := approvalService.VerifyApproveLogic(gormx, sqlx, verifyReq)
		if err != nil {
			return nil, err
		}

		for _, doc := range verifyRes.Documents {
			res = append(res, CreateSaleResponse{
				IsPass:           doc.IsPassPrice && doc.IsPassCredit && doc.IsPassInventory && doc.IsPassExpiryPrice,
				IsPassPrice:      doc.IsPassPrice,
				IsPassCredit:     doc.IsPassCredit,
				IsPassInventory:  doc.IsPassInventory,
				IsPassExpiryDate: doc.IsPassExpiryPrice,
				SaleCode:         doc.DocRef,
			})
		}
	}

	// check duplicate sale codes
	var existCount int64
	codes := make([]string, 0, len(createSales))
	for _, s := range createSales {
		codes = append(codes, s.SaleCode)
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
		if err := tx.Model(&models.Sale{}).
			Where("sale_code IN ?", codes).
			Count(&existCount).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		if existCount > 0 {
			tx.Rollback()
			return nil, errors.New("duplicate sale code detected")
		}
	}

	// Insert sales
	if len(createSales) > 0 {
		if err := tx.Create(&createSales).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// Insert sale items
	if len(createSaleItems) > 0 {
		if err := tx.Create(&createSaleItems).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return res, nil
}
