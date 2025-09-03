package models

import (
	"time"

	"github.com/google/uuid"
)

type Deposit struct {
	ID           uuid.UUID  `json:"id"`
	DepositCode  string     `json:"deposit_code"`
	DocRefType   string     `json:"doc_ref_type"`
	DocRef       string     `json:"doc_ref"`
	CustomerCode string     `json:"customer_code"`
	DepositDate  time.Time  `json:"deposit_date"` // วันที่รับมัดจำ
	AmountTotal  float64    `json:"amount_total"`
	AmountUsed   float64    `json:"amount_used"`
	Status       string     `json:"status"` // PENDING, COMPLETED, CANCELLED
	Remark       *string    `json:"remark"`
	CreateBy     string     `json:"create_by"`
	CreateDTM    time.Time  `json:"create_dtm"`
	UpdateBy     *string    `json:"update_by"`
	UpdateDTM    *time.Time `json:"update_dtm"`
}

func (Deposit) TableName() string { return "deposit" }
