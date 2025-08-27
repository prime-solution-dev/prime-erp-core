package approvalService

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/prime-solution-dev/prime-erp-core/Internal/db"
)

type VerifyApproveRequest struct {
	IsVerifyPriceUnit   bool       `json:"is_verify_price_unit"`
	IsVerifyPriceWeight bool       `json:"is_verify_price_weight"`
	IsVerifyCredit      bool       `json:"is_verify_credit"`
	IsVerifyExpiryPrice bool       `json:"is_verify_expiry_price"`
	IsVerifyInventory   bool       `json:"is_verify_inventory"`
	Documents           []document `json:"documents"`
}

type document struct {
	DocRef             string `json:"doc_ref"`
	CustomerCode       string `json:"customer_code"`
	EffectiveDatePrice string `json:"effective_date_price"`
	Items              []item `json:"items"`

	//Result
	IsPassPriceUnit   bool `json:"is_pass_price_unit"`
	IsPassPriceWeight bool `json:"is_pass_price_weight"`
	IsPassCredit      bool `json:"is_pass_credit"`
	IsPassExpiryPrice bool `json:"is_pass_expiry_price"`
	IsPassInventory   bool `json:"is_pass_inventory"`
}

type item struct {
	ItemRef           string  `json:"item_ref"`
	ProductCode       string  `json:"product_code"`
	Qty               float64 `json:"qty"`
	Unit              string  `json:"unit_code"`
	Weight            float64 `json:"weight"`
	PriceUnit         float64 `json:"price"`
	TransportCostUnit float64 `json:"transport_cost_unit"`
	TotalPrice        float64 `json:"total"`

	//Result
	IsPassPriceUnit   bool `json:"is_pass_price_unit"`
	IsPassPriceWeight bool `json:"is_pass_price_weight"`
	IsPassInventory   bool `json:"is_pass_inventory"`
}

type VerifyApproveResponse struct {
	IsPassVerify bool       `json:"is_pass_verify"`
	Documents    []document `json:"documents"`
}

func VerifyApprove(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	req := VerifyApproveRequest{}
	res := VerifyApproveResponse{}

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	sqlx, err := db.ConnectSqlx(`prime_erp_sale`)
	if err != nil {
		return nil, err
	}
	defer sqlx.Close()

	//TODO : vaerify logic from price service

	return res, nil
}
