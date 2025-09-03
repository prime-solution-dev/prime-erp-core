package approvalService

import (
	"fmt"
	"time"

	externalService "prime-erp-core/external/warehouse-service"
)

type VerifyInventoryRequest struct {
	CompanyCodes   string                   `json:"company_codes"`
	SiteCodes      string                   `json:"site_codes"`
	WarehouseCodes []string                 `json:"warehouse_codes"`
	StorageTypes   []string                 `json:"storage_types"`
	Products       []VerifyInventoryProduct `json:"products"`
	ToDate         *time.Time               `json:"to_date"`
}

type VerifyInventoryProduct struct {
	ProductCode string  `json:"product_code"`
	Qty         float64 `json:"qty"`
}

type VerifyInventoryResponse struct {
	IsPassInventory bool `json:"is_pass_inventory"`
}

func VerifyInventoryLogic(req VerifyInventoryRequest) (*VerifyInventoryResponse, error) {
	res := &VerifyInventoryResponse{}

	if len(req.Products) == 0 {
		return res, fmt.Errorf("require at least one product")
	}

	productExists := map[string]bool{}
	reqAtp := externalService.GetInventoryAtpRequest{
		CompanyCodes: []string{req.CompanyCodes},
		SiteCodes:    []string{req.SiteCodes},
		StorageTypes: req.StorageTypes,
		ToDate:       req.ToDate,
	}

	for _, p := range req.Products {
		if p.Qty <= 0 {
			continue
		}

		if !productExists[p.ProductCode] {
			reqAtp.ProductCodes = append(reqAtp.ProductCodes, p.ProductCode)
			productExists[p.ProductCode] = true
		}
	}

	resAtp, err := externalService.GetInventoryATP(reqAtp)
	if err != nil {
		return nil, err
	}

	remainingATP := map[string]float64{}
	for _, atp := range resAtp.ProductAtps {
		remainingATP[atp.ProductCode] += float64(atp.TodayAtpQty)
	}

	for _, p := range req.Products {
		if p.Qty <= 0 {
			continue
		}

		if remainingATP[p.ProductCode] >= p.Qty {
			remainingATP[p.ProductCode] -= p.Qty
		} else {
			res.IsPassInventory = false
			return res, nil
		}
	}

	res.IsPassInventory = true
	return res, nil
}
