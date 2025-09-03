package models

import (
	"time"

	"github.com/google/uuid"
)

type Credit struct {
	ID           uuid.UUID  `json:"id"`
	CustomerCode string     `json:"customer_code"`
	Amount       float64    `json:"amount"`
	EffectiveDTM time.Time  `json:"effective_dtm"`
	IsActive     bool       `json:"is_active"`
	DocRef       *string    `json:"doc_ref"`
	CreateBy     string     `json:"create_by"`
	CreateDTM    time.Time  `json:"create_dtm"`
	UpdateBy     *string    `json:"update_by"`
	UpdateDTM    *time.Time `json:"update_dtm"`

	Extras []CreditExtra `json:"extras"`
}

func (Credit) TableName() string { return "credit" }

type CreditExtra struct {
	ID           uuid.UUID  `json:"id"`
	CreditID     uuid.UUID  `json:"credit_id"`
	ExtraType    string     `json:"extra_type"`
	Amount       float64    `json:"amount"`
	EffectiveDTM time.Time  `json:"effective_dtm"`
	ExpireDTM    time.Time  `json:"expire_dtm"`
	DocRef       *string    `json:"doc_ref"`
	CreateBy     string     `json:"create_by"`
	CreateDTM    time.Time  `json:"create_dtm"`
	UpdateBy     *string    `json:"update_by"`
	UpdateDTM    *time.Time `json:"update_dtm"`
}

func (CreditExtra) TableName() string { return "credit_extra" }

type CreditTransaction struct {
	ID              uuid.UUID  `json:"id"`
	TransactionCode string     `json:"transaction_code"`
	TransactionType string     `json:"transaction_type"`
	Amount          float64    `json:"amount"`
	AdjustAmount    float64    `json:"adjust_amount"`
	EffectiveDTM    time.Time  `json:"effective_dtm"`
	ExpireDTM       *time.Time `json:"expire_dtm"`
	ForceExpireDTM  *time.Time `json:"force_expire_dtm"`
	IsApprove       bool       `json:"is_approve"`
	Status          string     `json:"status"`
	Reason          *string    `json:"reason"`
	CreateBy        string     `json:"create_by"`
	CreateDTM       time.Time  `json:"create_dtm"`
	UpdateBy        *string    `json:"update_by"`
	UpdateDTM       *time.Time `json:"update_dtm"`
}

func (CreditTransaction) TableName() string { return "credit_transaction" }
