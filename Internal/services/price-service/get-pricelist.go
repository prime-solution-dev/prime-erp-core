package priceService

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prime-solution-dev/prime-erp-core/Internal/db"
)

type GetPriceListGroupRequest struct {
	CompanyCode       string    `json:"company_code"`
	SiteCodes         []string  `json:"site_codes"`
	GroupCodes        []string  `json:"group_codes"`
	EffectiveDateFrom time.Time `json:"effective_date_from"`
	EffectiveDateTo   time.Time `json:"effective_date_to"`
	SubGroupCodes     []string  `json:"sub_group_codes"`
}

type GetPriceListGroupResponse struct {
	PriceListGroup
}

type PriceListGroup struct {
	ID                uuid.UUID             `json:"id"`
	CompanyCode       string                `json:"company_code"`
	SiteCode          string                `json:"site_code"`
	GroupCode         string                `json:"group_code"`
	PriceUnit         float64               `json:"price_unit"`
	PriceWeight       float64               `json:"price_weight"`
	BeforePriceUnit   float64               `json:"before_price_unit"`
	BeforePriceWeight float64               `json:"before_price_weight"`
	Currency          string                `json:"currency"`
	EffectiveDate     time.Time             `json:"effective_date"`
	ExtraPattern      string                `json:"extra_pattern"`
	Remark            string                `json:"remark"`
	Terms             []PriceListGroupTerm  `json:"terms"`
	Extras            []PriceListGroupExtra `json:"extras"`
	SubGroups         []SubGroup            `json:"sub_groups"`
}

type PriceListGroupTerm struct {
	ID         uuid.UUID `json:"id"`
	TermCode   string    `json:"term_code"`
	Pdc        float64   `json:"pdc"`
	PdcPercent float64   `json:"pdc_percent"`
	Due        float64   `json:"due"`
	DuePercent float64   `json:"due_percent"`
}

type PriceListGroupExtra struct {
	ID             uuid.UUID  `json:"id"`
	ExtraKey       string     `json:"extra_key"`
	LengthExtraKey int        `json:"length_extra_key"`
	Operator       string     `json:"operator"`
	CondRangeMin   float64    `json:"cond_range_min"`
	CondRangeMax   float64    `json:"cond_range_max"`
	GroupKeys      []GroupKey `json:"group_keys"`
}

type GroupKey struct {
	Code string `json:"code"`
	Seq  int    `json:"seq"`
}

type SubGroup struct {
	ID                        uuid.UUID  `json:"id"`
	SubGroupKey               string     `json:"subgroup_key"`
	IsTrading                 bool       `json:"is_trading"`
	PriceUnit                 float64    `json:"price_unit"`
	ExtraPriceUnit            float64    `json:"extra_price_unit"`
	TermPriceUnit             float64    `json:"term_price_unit"`
	TotalNetPriceUnit         float64    `json:"total_net_price_unit"`
	PriceWeight               float64    `json:"price_weight"`
	ExtraPriceWeight          float64    `json:"extra_price_weight"`
	TermPriceWeight           float64    `json:"term_price_weight"`
	TotalNetPriceWeight       float64    `json:"total_net_price_weight"`
	BeforePriceUnit           float64    `json:"before_price_unit"`
	BeforeExtraPriceUnit      float64    `json:"before_extra_price_unit"`
	BeforeTermPriceUnit       float64    `json:"before_term_price_unit"`
	BeforeTotalNetPriceUnit   float64    `json:"before_total_net_price_unit"`
	BeforePriceWeight         float64    `json:"before_price_weight"`
	BeforeExtraPriceWeight    float64    `json:"before_extra_price_weight"`
	BeforeTermPriceWeight     float64    `json:"before_term_price_weight"`
	BeforeTotalNetPriceWeight float64    `json:"before_total_net_price_weight"`
	EffectiveDate             time.Time  `json:"effective_date"`
	Remark                    string     `json:"remark"`
	GroupKeys                 []GroupKey `json:"group_keys"`
}

func GetPriceListGroup(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	var req GetPriceListGroupRequest
	var res []GetPriceListGroupResponse

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	user := `System`
	tenantId := uuid.New()

	_ = user
	_ = tenantId

	gormx, err := db.ConnectGORM(`prime_wms_warehouse`)
	if err != nil {
		return nil, err
	}
	defer db.CloseGORM(gormx)

	sqlx, err := db.ConnectSqlx(`prime_wms_warehouse`)
	if err != nil {
		return nil, err
	}
	defer sqlx.Close()

	return res, nil
}
