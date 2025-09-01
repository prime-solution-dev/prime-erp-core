package approvalService

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prime-solution-dev/prime-erp-core/Internal/db"
	priceService "github.com/prime-solution-dev/prime-erp-core/Internal/services/price-service"
)

type VerifyApproveRequest struct {
	IsVerifyPrice       bool     `json:"is_verify_price"`
	IsVerifyExpiryPrice bool     `json:"is_verify_expiry_price"`
	IsVerifyInventory   bool     `json:"is_verify_inventory"`
	IsVerifyCredit      bool     `json:"is_verify_credit"`
	Documents           document `json:"documents"`
}

type document struct {
	DocRef             string    `json:"doc_ref"`
	CustomerCode       string    `json:"customer_code"`
	EffectiveDatePrice time.Time `json:"effective_date_price"`
	Items              []item    `json:"items"`
}

type item struct {
	ItemRef                 string  `json:"item_ref"`
	ProductCode             string  `json:"product_code"`
	Qty                     float64 `json:"qty"`
	Unit                    string  `json:"unit_code"`
	Weight                  float64 `json:"weight"`
	PriceUnit               float64 `json:"price"`
	TotalPrice              float64 `json:"total"`
	TransportCostUnit       float64 `json:"transport_cost_unit"`
	TransportCostUnitWeight float64 `json:"transport_cost_unit_weight"`
}

type VerifyApproveResponse struct {
	IsPassPrice       bool     `json:"is_pass_price"`
	IsPassCredit      bool     `json:"is_pass_credit"`
	IsPassExpiryPrice bool     `json:"is_pass_expiry_price"`
	IsPassInventory   bool     `json:"is_pass_inventory"`
	Documents         document `json:"documents"`
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

	gormx, err := db.ConnectGORM(`prime_erp_sale`)
	if err != nil {
		return nil, err
	}
	defer db.CloseGORM(gormx)

	compareReq := priceService.GetComparePriceRequest{}

	for _, doc := range req.Documents.Items {
		//Price Verification
		item := priceService.ItemComparePrice{
			RefItem:     doc.ItemRef,
			ProductCode: doc.ProductCode,
			Qty:         doc.Qty,
			Unit:        doc.Unit,
			PriceUnit:   doc.PriceUnit,
			TotalPrice:  doc.TotalPrice,
			WeightUnit:  doc.Weight,
		}

		compareReq.Items = append(compareReq.Items, item)
		compareReq.TotalPrice += doc.TotalPrice
		compareReq.TotalWeight += doc.Weight
	}

	//Price Verification
	compareReq.UnitCode = `PCS`      //TODO: get from config
	compareReq.UnitCodeWeight = `KG` //TODO: get from config
	compareReq.TotalTransportCost = 0.0
	compareRes, err := verifyPrice(compareReq)
	if err != nil {
		return nil, err
	}

	if req.IsVerifyPrice && compareRes.IsPassPriceUnitAll == true && compareRes.IsPassPriceWeightAll == true {
		res.IsPassPrice = false
	}

	//Expiry Price Verification
	expPriceReq := VerifyExpiryPriceRequest{}
	expPriceReq.EffectiveDatePrice = req.Documents.EffectiveDatePrice
	expPrice, err := VerifyExpiryPriceLogic(gormx, expPriceReq)
	if err != nil {
		return nil, err
	}

	if req.IsVerifyExpiryPrice {
		res.IsPassExpiryPrice = expPrice.IsPassVerified
	}

	//TODO: Inventory Verification

	//TODO: Credit Verification

	return res, nil
}

func verifyPrice(compareReq priceService.GetComparePriceRequest) (*priceService.GetComparePriceResponse, error) {
	compareRes, err := priceService.ComparePrice(compareReq)
	if err != nil {
		return nil, err
	}

	return &compareRes, nil

}
