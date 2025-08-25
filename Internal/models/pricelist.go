package models

import (
	"time"

	"github.com/google/uuid"
)

// price_list_group
type PriceListGroup struct {
	ID                uuid.UUID `json:"id"`
	CompanyCode       string    `json:"company_code"`
	SiteCode          string    `json:"site_code"`
	GroupCode         string    `json:"group_code"`
	PriceUnit         float64   `json:"price_unit"`
	PriceWeight       float64   `json:"price_weight"`
	BeforePriceUnit   float64   `json:"before_price_unit"`
	BeforePriceWeight float64   `json:"before_price_weight"`
	Currency          string    `json:"currency"`
	EffectiveDate     time.Time `json:"effective_date"`
	ExtraPattern      string    `json:"extra_pattern"`
	Remark            string    `json:"remark"`
	CreateBy          string    `json:"create_by"`
	CreateDtm         time.Time `json:"create_dtm"`
	UpdateBy          string    `json:"update_by"`
	UpdateDtm         time.Time `json:"update_dtm"`
}

func (PriceListGroup) TableName() string { return "price_list_group" }

// price_list_group_history
type PriceListGroupHistory struct {
	ID                uuid.UUID `json:"id"`
	CompanyCode       string    `json:"company_code"`
	SiteCode          string    `json:"site_code"`
	GroupCode         string    `json:"group_code"`
	PriceUnit         float64   `json:"price_unit"`
	PriceWeight       float64   `json:"price_weight"`
	BeforePriceUnit   float64   `json:"before_price_unit"`
	BeforePriceWeight float64   `json:"before_price_weight"`
	Currency          string    `json:"currency"`
	EffectiveDate     time.Time `json:"effective_date"`
	ExpiryDate        time.Time `json:"expiry_date"`
	ExtraPattern      string    `json:"extra_pattern"`
	Remark            string    `json:"remark"`
	CreateBy          string    `json:"create_by"`
	CreateDtm         time.Time `json:"create_dtm"`
	UpdateBy          string    `json:"update_by"`
	UpdateDtm         time.Time `json:"update_dtm"`
}

func (PriceListGroupHistory) TableName() string { return "price_list_group_history" }

// price_list_group_term
type PriceListGroupTerm struct {
	ID               uuid.UUID `json:"id"`
	PriceListGroupID uuid.UUID `json:"price_list_group_id"`
	TermCode         string    `json:"term_code"`
	PDC              float64   `json:"pdc"`
	PDCPercent       int       `json:"pdc_percent"`
	Due              float64   `json:"due"`
	DuePercent       int       `json:"due_percent"`
	CreateBy         string    `json:"create_by"`
	CreateDtm        time.Time `json:"create_dtm"`
	UpdateBy         string    `json:"update_by"`
	UpdateDtm        time.Time `json:"update_dtm"`
}

func (PriceListGroupTerm) TableName() string { return "price_list_group_term" }

// price_list_group_extra
type PriceListGroupExtra struct {
	ID               uuid.UUID `json:"id"`
	PriceListGroupID uuid.UUID `json:"price_list_group_id"`
	ExtraKey         string    `json:"extra_key"`
	LengthExtraKey   int       `json:"length_extra_key"`
	Operator         string    `json:"operator"`
	CondRangeMin     float64   `json:"cond_range_min"`
	CondRangeMax     float64   `json:"cond_range_max"`
	CreateBy         string    `json:"create_by"`
	CreateDtm        time.Time `json:"create_dtm"`
	UpdateBy         string    `json:"update_by"`
	UpdateDtm        time.Time `json:"update_dtm"`
}

func (PriceListGroupExtra) TableName() string { return "price_list_group_extra" }

// group_extra_key
type GroupExtraKey struct {
	ID           uuid.UUID `json:"id"`
	GroupExtraID uuid.UUID `json:"group_extra_id"`
	Code         string    `json:"code"`
	Seq          int       `json:"seq"`
}

func (GroupExtraKey) TableName() string { return "group_extra_key" }

// price_list_sub_group
type PriceListSubGroup struct {
	ID                        uuid.UUID `json:"id"`
	PriceListGroupID          uuid.UUID `json:"price_list_group_id"`
	SubGroupKey               string    `json:"subgroup_key"`
	IsTrading                 bool      `json:"is_trading"`
	PriceUnit                 float64   `json:"price_unit"`
	ExtraPriceUnit            float64   `json:"extra_price_unit"`
	TermPriceUnit             float64   `json:"term_price_unit"`
	TotalNetPriceUnit         float64   `json:"total_net_price_unit"`
	PriceWeight               float64   `json:"price_weight"`
	ExtraPriceWeight          float64   `json:"extra_price_weight"`
	TermPriceWeight           float64   `json:"term_price_weight"`
	TotalNetPriceWeight       float64   `json:"total_net_price_weight"`
	BeforePriceUnit           float64   `json:"before_price_unit"`
	BeforeExtraPriceUnit      float64   `json:"before_extra_price_unit"`
	BeforeTermPriceUnit       float64   `json:"before_term_price_unit"`
	BeforeTotalNetPriceUnit   float64   `json:"before_total_net_price_unit"`
	BeforePriceWeight         float64   `json:"before_price_weight"`
	BeforeExtraPriceWeight    float64   `json:"before_extra_price_weight"`
	BeforeTermPriceWeight     float64   `json:"before_term_price_weight"`
	BeforeTotalNetPriceWeight float64   `json:"before_total_net_price_weight"`
	EffectiveDate             time.Time `json:"effective_date"`
	Remark                    string    `json:"remark"`
	CreateBy                  string    `json:"create_by"`
	CreateDtm                 time.Time `json:"create_dtm"`
	UpdateBy                  string    `json:"update_by"`
	UpdateDtm                 time.Time `json:"update_dtm"`
}

func (PriceListSubGroup) TableName() string { return "price_list_sub_group" }

// sub_group_key
type SubGroupKey struct {
	ID         uuid.UUID `json:"id"`
	SubGroupID uuid.UUID `json:"sub_group_id"`
	Code       string    `json:"code"`
	Seq        int       `json:"seq"`
}

func (SubGroupKey) TableName() string { return "sub_group_key" }

// price_list_sub_group_history
type PriceListSubGroupHistory struct {
	ID                        uuid.UUID `json:"id"`
	PriceListGroupID          uuid.UUID `json:"price_list_group_id"`
	SubGroupKey               string    `json:"subgroup_key"`
	IsTrading                 bool      `json:"is_trading"`
	PriceUnit                 float64   `json:"price_unit"`
	ExtraPriceUnit            float64   `json:"extra_price_unit"`
	TermPriceUnit             float64   `json:"term_price_unit"`
	TotalNetPriceUnit         float64   `json:"total_net_price_unit"`
	PriceWeight               float64   `json:"price_weight"`
	ExtraPriceWeight          float64   `json:"extra_price_weight"`
	TermPriceWeight           float64   `json:"term_price_weight"`
	TotalNetPriceWeight       float64   `json:"total_net_price_weight"`
	BeforePriceUnit           float64   `json:"before_price_unit"`
	BeforeExtraPriceUnit      float64   `json:"before_extra_price_unit"`
	BeforeTermPriceUnit       float64   `json:"before_term_price_unit"`
	BeforeTotalNetPriceUnit   float64   `json:"before_total_net_price_unit"`
	BeforePriceWeight         float64   `json:"before_price_weight"`
	BeforeExtraPriceWeight    float64   `json:"before_extra_price_weight"`
	BeforeTermPriceWeight     float64   `json:"before_term_price_weight"`
	BeforeTotalNetPriceWeight float64   `json:"before_total_net_price_weight"`
	EffectiveDate             time.Time `json:"effective_date"`
	ExpiryDate                time.Time `json:"expiry_date"`
	Remark                    string    `json:"remark"`
	CreateBy                  string    `json:"create_by"`
	CreateDtm                 time.Time `json:"create_dtm"`
	UpdateBy                  string    `json:"update_by"`
	UpdateDtm                 time.Time `json:"update_dtm"`
}

func (PriceListSubGroupHistory) TableName() string { return "price_list_sub_group_history" }

// sub_group_key_history
type SubGroupKeyHistory struct {
	ID                uuid.UUID `json:"id"`
	SubGroupHistoryID uuid.UUID `json:"sub_group_history_id"`
	Code              string    `json:"code"`
	Seq               int       `json:"seq"`
}

func (SubGroupKeyHistory) TableName() string { return "sub_group_key_history" }

// payment_term
type PaymentTerm struct {
	ID        uuid.UUID `json:"id"`
	TermCode  string    `json:"term_code"`
	TermName  string    `json:"term_name"`
	TermType  string    `json:"term_type"`
	CreateBy  string    `json:"create_by"`
	CreateDtm time.Time `json:"create_dtm"`
	UpdateBy  string    `json:"update_by"`
	UpdateDtm time.Time `json:"update_dtm"`
}

func (PaymentTerm) TableName() string { return "payment_term" }

// extra_pattern
type ExtraPattern struct {
	PatternCode string `json:"pattern_code"`
	PatternName string `json:"pattern_name"`
}

func (ExtraPattern) TableName() string { return "extra_pattern" }
