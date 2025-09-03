package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Approval struct {
	ID             uuid.UUID       `json:"id"`
	ApproveCode    string          `json:"approve_code"`
	ApproveTopic   string          `json:"approve_topic"`
	DocumentType   string          `json:"document_type"`
	DocumentCode   string          `json:"document_code"`
	DocumentData   json.RawMessage `json:"document_data"`
	ActionDate     *time.Time      `json:"action_date"`
	Status         string          `json:"status"`
	Remark         *string         `json:"remark"`
	CurrentStepSeq int             `json:"current_step_seq"`
	CreateDate     time.Time       `json:"create_date"`
	CreateBy       string          `json:"create_by"`
	UpdateDate     *time.Time      `json:"update_date"`
	UpdateBy       *string         `json:"update_by"`
}

func (Approval) TableName() string { return "approval" }

type ApprovalItem struct {
	ID          uuid.UUID       `json:"id"`
	ApprovalID  uuid.UUID       `json:"approval_id"`
	StepSeq     int             `json:"step_seq"`
	IsCondition bool            `json:"is_condition"`
	Condition   json.RawMessage `json:"condition"`
	Status      string          `json:"status"`
	ActionBy    *string         `json:"action_by"`
	ActionDate  *time.Time      `json:"action_date"`
	CreateDate  time.Time       `json:"create_date"`
	CreateBy    string          `json:"create_by"`
	UpdateDate  *time.Time      `json:"update_date"`
	UpdateBy    *string         `json:"update_by"`
}

func (ApprovalItem) TableName() string { return "approval_item" }

type ApprovalItemPermission struct {
	ID             uuid.UUID `json:"id"`
	ApprovalItemID uuid.UUID `json:"approval_item_id"`
	UserCode       string    `json:"user_code"`
}

func (ApprovalItemPermission) TableName() string { return "approval_item_permission" }
