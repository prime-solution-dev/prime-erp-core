package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- QUOTATION --------------------

type Quotation struct {
	ID                  uuid.UUID  `json:"id"`
	QuotationCode       string     `json:"quotation_code"`
	CustomerCode        string     `json:"customer_code"`
	CustomerName        string     `json:"customer_name"`
	DeliveryDate        *time.Time `json:"delivery_date"`
	SoldToCode          string     `json:"sold_to_code"`
	SoldToAddress       string     `json:"sold_to_address"`
	BillToCode          string     `json:"bill_to_code"`
	BillToAddress       string     `json:"bill_to_address"`
	ShipToCode          string     `json:"ship_to_code"`
	ShipToType          string     `json:"ship_to_type"`
	ShipToAddress       string     `json:"ship_to_address"`
	DeliveryMethod      string     `json:"delivery_method"`
	TransportCostType   string     `json:"transport_cost_type"`
	TotalTransportCost  float64    `json:"total_transport_cost"`
	TotalPrice          float64    `json:"total_price"`
	TotalWeight         float64    `json:"total_weight"`
	TotalNetPrice       float64    `json:"total_net_price"`
	TotalNetPriceWeight float64    `json:"total_net_price_weight"`
	PassPrice           bool       `json:"pass_price"`
	PaymentMethod       string     `json:"payment_method"`
	PaymentTermCode     string     `json:"payment_term_code"`
	SalePersonCode      string     `json:"sale_person_code"`
	EffectiveDatePrice  *time.Time `json:"effective_date_price"`
	ExpirePriceDay      int        `json:"expire_price_day"`
	ExpirePriceDate     *time.Time `json:"expire_price_date"`
	PassPriceList       string     `json:"pass_price_list"`
	PassAtpCheck        string     `json:"pass_atp_check"`
	PassCreditLimit     string     `json:"pass_credit_limit"`
	PassPriceExpire     string     `json:"pass_price_expire"`
	Status              string     `json:"status"`
	Remark              string     `json:"remark"`
	IsApproved          bool       `json:"is_approved"`
	StatusApprove       string     `json:"status_approve"`
	CreateDate          *time.Time `json:"create_date"`
	CreateBy            string     `json:"create_by"`
	UpdateDate          *time.Time `json:"update_date"`
	UpdateBy            string     `json:"update_by"`
}

func (Quotation) TableName() string { return "quotation" }

// -------------------- QUOTATION ITEM --------------------

type QuotationItem struct {
	ID                      uuid.UUID  `json:"id"`
	QuotationID             uuid.UUID  `json:"quotation_id"`
	QuotationItem           string     `json:"quotation_item"`
	ProductCode             string     `json:"product_code"`
	ProductDesc             string     `json:"product_desc"`
	Qty                     float64    `json:"qty"`
	Unit                    string     `json:"unit"`
	PriceListUnit           float64    `json:"price_list_unit"`
	SaleQty                 float64    `json:"sale_qty"`
	SaleUnit                string     `json:"sale_unit"`
	SaleUnitType            string     `json:"sale_unit_type"`
	PassPrice               string     `json:"pass_price"`
	PassWeight              string     `json:"pass_weight"`
	PriceUnit               float64    `json:"price_unit"`
	TotalPrice              float64    `json:"total_price"`
	TransportCostUnit       float64    `json:"transport_cost_unit"`
	TotalNetPrice           float64    `json:"total_net_price"`
	NetPriceUnit            float64    `json:"net_price_unit"`
	WeightUnit              float64    `json:"weight_unit"`
	AvgWeightUnit           float64    `json:"avg_weight_unit"`
	TotalWeight             float64    `json:"total_weight"`
	TransportCostWeightUnit float64    `json:"transport_cost_weight_unit"`
	TotalNetPriceWeight     float64    `json:"total_net_price_weight"`
	NetPriceWeightUnit      float64    `json:"net_price_weight_unit"`
	Status                  string     `json:"status"`
	Remark                  string     `json:"remark"`
	CreateDate              *time.Time `json:"create_date"`
	CreateBy                string     `json:"create_by"`
	UpdateDate              *time.Time `json:"update_date"`
	UpdateBy                string     `json:"update_by"`
}

func (QuotationItem) TableName() string { return "quotation_item" }

// -------------------- SALE --------------------

type Sale struct {
	ID                  uuid.UUID  `json:"id"`
	SaleCode            string     `json:"sale_code"`
	CustomerCode        string     `json:"customer_code"`
	CustomerName        string     `json:"customer_name"`
	DeliveryDate        *time.Time `json:"delivery_date"`
	SoldToCode          string     `json:"sold_to_code"`
	SoldToAddress       string     `json:"sold_to_address"`
	BillToCode          string     `json:"bill_to_code"`
	BillToAddress       string     `json:"bill_to_address"`
	ShipToCode          string     `json:"ship_to_code"`
	ShipToAddress       string     `json:"ship_to_address"`
	ShipToType          string     `json:"ship_to_type"`
	DeliveryMethod      string     `json:"delivery_method"`
	TransportCostType   string     `json:"transport_cost_type"`
	TransportCost       float64    `json:"transport_cost"`
	TotalPrice          float64    `json:"total_price"`
	TotalWeight         float64    `json:"total_weight"`
	TotalNetPrice       float64    `json:"total_net_price"`
	TotalNetPriceWeight float64    `json:"total_net_price_weight"`
	PaymentMethod       string     `json:"payment_method"`
	PaymentTermCode     string     `json:"payment_term_code"`
	SalePersonCode      string     `json:"sale_person_code"`
	EffectiveDatePrice  *time.Time `json:"effective_date_price"`
	ExpirePriceDay      int        `json:"expire_price_day"`
	ExpirePriceDate     *time.Time `json:"expire_price_date"`
	PassPriceList       string     `json:"pass_price_list"`
	PassAtpCheck        string     `json:"pass_atp_check"`
	PassCreditLimit     string     `json:"pass_credit_limit"`
	PassPriceExpire     string     `json:"pass_price_expire"`
	Status              string     `json:"status"`
	Remark              string     `json:"remark"`
	IsApproved          bool       `json:"is_approved"`
	StatusApprove       string     `json:"status_approve"`
	CreateDate          *time.Time `json:"create_date"`
	CreateBy            string     `json:"create_by"`
	UpdateDate          *time.Time `json:"update_date"`
	UpdateBy            string     `json:"update_by"`
}

func (Sale) TableName() string { return "sale" }

// -------------------- SALE ITEM --------------------

type SaleItem struct {
	ID                      uuid.UUID  `json:"id"`
	SaleID                  uuid.UUID  `json:"sale_id"`
	SaleItem                string     `json:"sale_item"`
	ProductCode             string     `json:"product_code"`
	ProductDesc             string     `json:"product_desc"`
	Qty                     float64    `json:"qty"`
	OriginQty               float64    `json:"origin_qty"`
	Unit                    string     `json:"unit"`
	PriceListUnit           float64    `json:"price_list_unit"`
	SaleQty                 float64    `json:"sale_qty"`
	SaleUnit                string     `json:"sale_unit"`
	SaleUnitType            string     `json:"sale_unit_type"`
	PassPriceUnit           string     `json:"pass_price_unit"`
	PassPriceWeight         string     `json:"pass_price_weight"`
	PriceUnit               float64    `json:"price_unit"`
	TotalPrice              float64    `json:"total_price"`
	TransportCostUnit       float64    `json:"transport_cost_unit"`
	TotalNetPrice           float64    `json:"total_net_price"`
	NetPriceUnit            float64    `json:"net_price_unit"`
	WeightUnit              float64    `json:"weight_unit"`
	AvgWeightUnit           float64    `json:"avg_weight_unit"`
	TotalWeight             float64    `json:"total_weight"`
	TransportCostWeightUnit float64    `json:"transport_cost_weight_unit"`
	TotalNetPriceWeight     float64    `json:"total_net_price_weight"`
	NetPriceWeightUnit      float64    `json:"net_price_weight_unit"`
	Status                  string     `json:"status"`
	Remark                  string     `json:"remark"`
	CreateDate              *time.Time `json:"create_date"`
	CreateBy                string     `json:"create_by"`
	UpdateDate              *time.Time `json:"update_date"`
	UpdateBy                string     `json:"update_by"`
}

func (SaleItem) TableName() string { return "sale_item" }
