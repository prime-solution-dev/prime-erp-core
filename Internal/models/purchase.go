package models

import "time"

type PrePurchase struct {
	ID                          string    `json:"id"`
	PurchaseCode                string    `json:"purchase_code"`
	PurchaseType                string    `json:"purchase_type"`
	CompanyCode                 string    `json:"company_code"`
	SiteCode                    string    `json:"site_code"`
	DocRefType                  string    `json:"doc_ref_type"`
	DocRef                      string    `json:"doc_ref"`
	SupplierCode                string    `json:"supplier_code"`
	DeliveryAddress             string    `json:"delivery_address"`
	Status                      string    `json:"status"`
	TotalAmount                 float64   `json:"total_amount"`
	TotalWeight                 float64   `json:"total_weight"`
	SubtotalExclTransport       float64   `json:"subtotal_excl_transport"`
	SubtotalWeightExclTransport float64   `json:"subtotal_weight_excl_transport"`
	IsApproved                  bool      `json:"is_approved"`
	StatusApprove               string    `json:"status_approve"`
	Remark                      string    `json:"remark"`
	CreateBy                    string    `json:"create_by"`
	CreateDtm                   time.Time `json:"create_dtm"`
	UpdateBy                    string    `json:"update_by"`
	UpdateDate                  time.Time `json:"update_date"`
}

func (PrePurchase) TableName() string {
	return "pre_purchase"
}

type PrePurchaseItem struct {
	ID                             string    `json:"id"`
	PurchaseID                     string    `json:"purchase_id"`
	PreItem                        string    `json:"pre_item"`
	HierarchyType                  string    `json:"hierarchy_type"`
	HierarchyCode                  string    `json:"hierarchy_code"`
	DocRefItem                     string    `json:"doc_ref_item"`
	Qty                            float64   `json:"qty"`
	Unit                           string    `json:"unit"`
	PurchaseQty                    float64   `json:"purchase_qty"`
	PurchaseUnit                   string    `json:"purchase_unit"`
	PurchaseUnitType               string    `json:"purchase_unit_type"`
	PriceUnit                      float64   `json:"price_unit"`
	TotalDiscount                  float64   `json:"total_discount"`
	TotalAmount                    float64   `json:"total_amount"`
	TransportCostUnit              float64   `json:"transport_cost_unit"`
	SubtotalExclTransport          float64   `json:"subtotal_excl_transport"`
	NetPriceUnitExclTransport      float64   `json:"net_price_unit_excl_transport"`
	WeightUnit                     float64   `json:"weight_unit"`
	AvgWeightUnit                  float64   `json:"avg_weight_unit"`
	TotalWeight                    float64   `json:"total_weight"`
	TransportCostWeightUnit        float64   `json:"transport_cost_weight_unit"`
	SubtotalWeightExclTransport    float64   `json:"subtotal_weight_excl_transport"`
	NetPricePerWeightExclTransport float64   `json:"net_price_per_weight_excl_transport"`
	Status                         string    `json:"status"`
	Remark                         string    `json:"remark"`
	CreateDate                     time.Time `json:"create_date"`
	CreateBy                       string    `json:"create_by"`
	UpdateDate                     time.Time `json:"update_date"`
	UpdateBy                       string    `json:"update_by"`
}

func (PrePurchaseItem) TableName() string {
	return "pre_purchase_item"
}

type Purchase struct {
	ID                          string    `json:"id"`
	PurchaseCode                string    `json:"purchase_code"`
	PurchaseType                string    `json:"purchase_type"`
	CompanyCode                 string    `json:"company_code"`
	SiteCode                    string    `json:"site_code"`
	DocRefType                  string    `json:"doc_ref_type"`
	DocRef                      string    `json:"doc_ref"`
	SupplierCode                string    `json:"supplier_code"`
	DeliveryAddress             string    `json:"delivery_address"`
	Status                      string    `json:"status"`
	TotalAmount                 float64   `json:"total_amount"`
	TotalWeight                 float64   `json:"total_weight"`
	SubtotalExclTransport       float64   `json:"subtotal_excl_transport"`
	SubtotalWeightExclTransport float64   `json:"subtotal_weight_excl_transport"`
	IsApproved                  bool      `json:"is_approved"`
	StatusApprove               string    `json:"status_approve"`
	Remark                      string    `json:"remark"`
	CreateBy                    string    `json:"create_by"`
	CreateDtm                   time.Time `json:"create_dtm"`
	UpdateBy                    string    `json:"update_by"`
	UpdateDate                  time.Time `json:"update_date"`
}

func (Purchase) TableName() string {
	return "purchase"
}

type PurchaseItem struct {
	ID                             string    `json:"id"`
	PurchaseID                     string    `json:"purchase_id"`
	PurchaseItem                   string    `json:"purchase_item"`
	ProductCode                    string    `json:"product_code"`
	DocRefItem                     string    `json:"doc_ref_item"`
	Qty                            float64   `json:"qty"`
	Unit                           string    `json:"unit"`
	PurchaseQty                    float64   `json:"purchase_qty"`
	PurchaseUnit                   string    `json:"purchase_unit"`
	PurchaseUnitType               string    `json:"purchase_unit_type"`
	PriceUnit                      float64   `json:"price_unit"`
	TotalDiscount                  float64   `json:"total_discount"`
	TotalAmount                    float64   `json:"total_amount"`
	TransportCostUnit              float64   `json:"transport_cost_unit"`
	SubtotalExclTransport          float64   `json:"subtotal_excl_transport"`
	NetPriceUnitExclTransport      float64   `json:"net_price_unit_excl_transport"`
	WeightUnit                     float64   `json:"weight_unit"`
	AvgWeightUnit                  float64   `json:"avg_weight_unit"`
	TotalWeight                    float64   `json:"total_weight"`
	TransportCostWeightUnit        float64   `json:"transport_cost_weight_unit"`
	SubtotalWeightExclTransport    float64   `json:"subtotal_weight_excl_transport"`
	NetPricePerWeightExclTransport float64   `json:"net_price_per_weight_excl_transport"`
	Status                         string    `json:"status"`
	Remark                         string    `json:"remark"`
	CreateDate                     time.Time `json:"create_date"`
	CreateBy                       string    `json:"create_by"`
	UpdateDate                     time.Time `json:"update_date"`
	UpdateBy                       string    `json:"update_by"`
}

func (PurchaseItem) TableName() string {
	return "purchase_item"
}
