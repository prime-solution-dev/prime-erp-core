package quotationService

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prime-solution-dev/prime-erp-core/Internal/db"
	"github.com/prime-solution-dev/prime-erp-core/Internal/models"
)

type CreateQuotationRequest struct {
	IsVerifyPrice bool               `json:"is_verify_price"` // true = verify, if not verified can't create
	Quotations    []QuotationPayload `json:"quotations"`
}

type QuotationPayload struct {
	models.Quotation
	Items []models.QuotationItem `json:"items"`
}

type CreateQuotationResponse struct {
	IsVerified    bool   `json:"is_verified"`
	QuotationCode string `json:"quotation_code"`
}

func CreateQuotation(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	req := CreateQuotationRequest{}
	res := []CreateQuotationResponse{}

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	gormx, err := db.ConnectGORM(`prime_erp_sale`)
	if err != nil {
		return nil, err
	}
	defer db.CloseGORM(gormx)

	user := `system` // TODO: get from ctx
	now := time.Now()

	tx := gormx.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	createQuotations := []models.Quotation{}
	createQuotationItems := []models.QuotationItem{}

	for _, quotationReq := range req.Quotations {
		tempQuotation := quotationReq.Quotation
		tempQuotation.ID = uuid.New()

		if tempQuotation.QuotationCode == "" {
			tempQuotation.QuotationCode = uuid.New().String()
		}

		tempQuotation.CreateDate = &now
		tempQuotation.CreateBy = user
		tempQuotation.UpdateDate = &now
		tempQuotation.UpdateBy = user

		createQuotations = append(createQuotations, tempQuotation)

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
		}
	}

	//TODO: implement price verification

	// check duplicate quotation codes
	var existCount int64
	codes := make([]string, 0, len(createQuotations))
	for _, q := range createQuotations {
		codes = append(codes, q.QuotationCode)
	}

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
