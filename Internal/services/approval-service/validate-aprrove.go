package approvalService

import (
	"encoding/json"
	"errors"
	"time"

	"prime-erp-core/Internal/db"
	priceService "prime-erp-core/Internal/services/price-service"

	"github.com/gin-gonic/gin"
)

type VerifyApproveRequest struct {
	IsVerifyPrice       bool      `json:"is_verify_price"`
	IsVerifyExpiryPrice bool      `json:"is_verify_expiry_price"`
	IsVerifyInventory   bool      `json:"is_verify_inventory"`
	IsVerifyCredit      bool      `json:"is_verify_credit"`
	SaleDate            time.Time `json:"sale_date"`
	CompanyCode         string    `json:"company_code"`
	SiteCode            string    `json:"site_code"`
	WarehouseCode       []string  `json:"warehouse_code"`
	StorageType         []string  `json:"storage_type"`
	Documents           document  `json:"documents"`
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
	isVerifyWithOldTransportCost := false

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	if len(req.Documents.DocRef) == 0 {
		return nil, errors.New("document reference is required")
	}

	if len(req.Documents.Items) == 0 {
		return nil, errors.New("require at least one item")
	}

	if req.SaleDate.IsZero() {
		return nil, errors.New("sale date is required")
	}

	if req.CompanyCode == `` {
		return nil, errors.New("company code is required")
	}

	if req.SiteCode == `` {
		return nil, errors.New("site code is required")
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

		//Optional for transport cost verification
		if isVerifyWithOldTransportCost {
			item.TransportCostUnit = &doc.TransportCostUnit
			item.TransportCostUnitWeight = &doc.TransportCostUnitWeight
		}

		compareReq.Items = append(compareReq.Items, item)
		compareReq.TotalPrice += doc.TotalPrice
		compareReq.TotalWeight += doc.Weight
	}

	//Price Verification
	if req.IsVerifyPrice {
		compareReq.UnitCode = `PCS`      //TODO: get from config
		compareReq.UnitCodeWeight = `KG` //TODO: get from config

		compareReq.TotalTransportCost = 0.0
		compareRes, err := verifyPrice(compareReq)
		if err != nil {
			return nil, err
		}

		if compareRes.IsPassPriceUnitAll == true && compareRes.IsPassPriceWeightAll == true {
			res.IsPassPrice = false
		}
	}

	//Expiry Price Verification
	if req.IsVerifyPrice {
		expPriceReq := VerifyExpiryPriceRequest{}
		expPriceReq.EffectiveDatePrice = req.Documents.EffectiveDatePrice
		expPrice, err := VerifyExpiryPriceLogic(gormx, expPriceReq)
		if err != nil {
			return nil, err
		}

		res.IsPassExpiryPrice = expPrice.IsPassVerified
	}

	//Inventory Verification
	if req.IsVerifyInventory {
		invReq := VerifyInventoryRequest{}
		invReq.CompanyCodes = req.CompanyCode
		invReq.SiteCodes = req.SiteCode
		invReq.StorageTypes = req.StorageType
		invReq.WarehouseCodes = req.WarehouseCode
		invReq.ToDate = &req.SaleDate

		for _, doc := range req.Documents.Items {
			product := VerifyInventoryProduct{
				ProductCode: doc.ProductCode,
				Qty:         doc.Qty,
			}

			invReq.Products = append(invReq.Products, product)
		}

		invRes, err := VerifyInventoryLogic(invReq)
		if err != nil {
			return nil, err
		}

		res.IsPassInventory = invRes.IsPassInventory
	}

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
