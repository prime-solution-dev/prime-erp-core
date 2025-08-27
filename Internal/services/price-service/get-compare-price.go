package priceService

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/gin-gonic/gin"
)

type GetComparePriceRequest struct {
	TotalPrice         float64            `json:"total_price"`
	TotalWeight        float64            `json:"total_weight"`
	TotalTransportCost float64            `json:"transport_cost"`
	UnitCode           string             `json:"unit_code"`        // PCS
	UnitCodeWeight     string             `json:"unit_code_weight"` //KG
	Items              []ItemComparePrice `json:"items"`
}

type GetComparePriceResponse struct {
	TotalPrice           float64            `json:"total_price"`
	TotalWeight          float64            `json:"total_weight"`
	TotalTransportCost   float64            `json:"total_transport_cost"`
	TotalNetPrice        float64            `json:"total_net_price"`
	AvgNetPriceUnit      float64            `json:"avg_net_price_unit"`
	AvgNetPriceWeight    float64            `json:"avg_net_price_weight"`
	IsPassPriceUnitAll   bool               `json:"is_pass_price_unit_all"`
	IsPassPriceWeightAll bool               `json:"is_pass_price_weight_all"`
	Items                []ItemComparePrice `json:"items"`
}

type ItemComparePrice struct {
	RefItem       string  `json:"ref_item"`
	ProductCode   string  `json:"product_code"`
	Qty           float64 `json:"qty"`
	Unit          string  `json:"unit_code"`
	SaleUnit      string  `json:"sale_unit"`
	SaleUnitType  string  `json:"sale_unit_type"`
	PriceListUnit float64 `json:"price_list"`
	PriceUnit     float64 `json:"price"`
	TotalPrice    float64 `json:"total_price"`
	WeightUnit    float64 `json:"weight_unit"`
	AvgWeightUnit float64 `json:"avg_weight_unit"`
	TotalWeight   float64 `json:"total_weight"`

	// Results
	TransportCostUnit       float64 `json:"transport_cost_unit"`
	TotalNetPrice           float64 `json:"total_net_price"`
	NetPriceUnit            float64 `json:"net_price_unit"`
	PriceDiffUnit           float64 `json:"price_diff_unit"`
	IsPassPriceUnit         bool    `json:"is_pass_price_unit"`
	TransportCostUnitWeight float64 `json:"transport_cost_unit_weight"`
	TotalNetPriceWeight     float64 `json:"total_net_price_weight"`
	NetPriceUnitWeight      float64 `json:"net_price_unit_weight"`
	PriceDiffUnitWeight     float64 `json:"price_diff_unit_weight"`
	IsPassPriceWeight       bool    `json:"is_pass_price_weight"`
}

func GetComparePrice(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	req := GetComparePriceRequest{}

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	return ComparePrice(req)
}

func ComparePrice(req GetComparePriceRequest) (GetComparePriceResponse, error) {
	res := GetComparePriceResponse{}

	if len(req.Items) == 0 {
		return GetComparePriceResponse{}, fmt.Errorf("items is required")
	}

	totalPriceAll := req.TotalPrice
	totalWeightAll := req.TotalWeight
	totalTransportCostAll := req.TotalTransportCost
	unitCode := req.UnitCode
	unitCodeWeight := req.UnitCodeWeight

	sumTransportCostUnit := 0.0
	sumTransportCostUnitWeight := 0.0
	for _, item := range req.Items {
		newItem := item

		// Transport Cost Unit
		if totalPriceAll > 0 {
			newItem.TransportCostUnit = round2(item.TotalPrice / totalPriceAll * totalTransportCostAll)
		} else {
			newItem.TransportCostUnit = 0
		}
		sumTransportCostUnit += newItem.TransportCostUnit

		// Transport Cost Unit Weight
		if totalWeightAll > 0 {
			newItem.TransportCostUnitWeight = round2(item.TotalWeight / totalWeightAll * totalTransportCostAll)
		} else {
			newItem.TransportCostUnitWeight = 0
		}
		sumTransportCostUnitWeight += newItem.TransportCostUnitWeight

		res.Items = append(res.Items, newItem)
	}

	// Sort by total price
	sort.SliceStable(res.Items, func(i, j int) bool {
		return res.Items[i].TotalPrice > res.Items[j].TotalPrice
	})

	// Adjust Transport Cost Unit
	diffTransportCostUnit := totalTransportCostAll - sumTransportCostUnit
	diffTransportCostUnitWeight := totalTransportCostAll - sumTransportCostUnitWeight

	for i := range res.Items {
		if diffTransportCostUnit == 0 {
			break
		}
		if diffTransportCostUnit > 0 {
			res.Items[i].TransportCostUnit += 0.01
			diffTransportCostUnit -= 0.01
		} else if diffTransportCostUnit < 0 {
			res.Items[i].TransportCostUnit -= 0.01
			diffTransportCostUnit += 0.01
		}
	}

	for i := range res.Items {
		if diffTransportCostUnitWeight == 0 {
			break
		}
		if diffTransportCostUnitWeight > 0 {
			res.Items[i].TransportCostUnitWeight += 0.01
			diffTransportCostUnitWeight -= 0.01
		} else if diffTransportCostUnitWeight < 0 {
			res.Items[i].TransportCostUnitWeight -= 0.01
			diffTransportCostUnitWeight += 0.01
		}
	}

	// Calculate Net Price Unit and Net Price Weight
	var totalNetPrice, totalNetPriceWeight float64
	passUnitAll := true
	passWeightAll := true

	for i := range res.Items {
		item := res.Items[i]

		// ---- Unit ----
		item.TotalNetPrice = round2(item.TotalPrice - item.TransportCostUnit)

		if item.SaleUnit == unitCode && item.Qty > 0 {
			item.NetPriceUnit = round2(item.TotalNetPrice / item.Qty)
		} else if item.SaleUnit == unitCodeWeight && item.TotalWeight > 0 {
			item.NetPriceUnit = round2(item.TotalNetPrice / item.TotalWeight)
		} else {
			item.NetPriceUnit = 0
		}

		item.PriceDiffUnit = round2(item.NetPriceUnit - item.PriceListUnit)
		item.IsPassPriceUnit = item.PriceDiffUnit >= 0
		if !item.IsPassPriceUnit {
			passUnitAll = false
		}

		// ---- Weight ----
		item.TotalNetPriceWeight = round2(item.TotalPrice - item.TransportCostUnitWeight)

		if item.SaleUnit == unitCode && item.Qty > 0 {
			item.NetPriceUnitWeight = round2(item.TotalNetPriceWeight / item.Qty)
		} else if item.SaleUnit == unitCodeWeight && item.TotalWeight > 0 {
			item.NetPriceUnitWeight = round2(item.TotalNetPriceWeight / item.TotalWeight)
		} else {
			item.NetPriceUnitWeight = 0
		}

		item.PriceDiffUnitWeight = round2(item.NetPriceUnitWeight - item.PriceListUnit)
		item.IsPassPriceWeight = item.PriceDiffUnitWeight >= 0
		if !item.IsPassPriceWeight {
			passWeightAll = false
		}

		// update item to response
		res.Items[i] = item

		totalNetPrice += item.TotalNetPrice
		totalNetPriceWeight += item.TotalNetPriceWeight
	}

	// summary response
	res.TotalPrice = round2(totalPriceAll)
	res.TotalWeight = round2(totalWeightAll)
	res.TotalTransportCost = round2(totalTransportCostAll)
	res.TotalNetPrice = round2(totalNetPrice)

	if totalWeightAll > 0 {
		res.AvgNetPriceWeight = round2(totalNetPriceWeight / totalWeightAll)
	}

	if totalPriceAll > 0 {
		res.AvgNetPriceUnit = round2(totalNetPrice / totalPriceAll)
	}

	res.IsPassPriceUnitAll = passUnitAll
	res.IsPassPriceWeightAll = passWeightAll

	return res, nil
}

func round2(val float64) float64 {
	return math.Round(val*100) / 100
}
