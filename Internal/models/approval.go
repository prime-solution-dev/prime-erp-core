package models

import (
	"time"

	"github.com/google/uuid"
)

type Approval struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;not null" json:"id"`
	ApproveCode   string         `gorm:"type:varchar(50);not null" json:"approve_code"`
	ApproveTopic  string         `gorm:"type:varchar(255);not null" json:"approve_topic"`
	DocumentType  string         `gorm:"type:varchar(50);not null" json:"document_type"`
	DocumentCode  string         `gorm:"type:varchar(50);not null" json:"document_code"`
	DocumentData  string         `gorm:"type:jsonb;not null" json:"document_data"`
	ActionDate    time.Time      `gorm:"type:timestamp;not null" json:"action_date"`
	Status        string         `gorm:"type:varchar(20);not null" json:"status"`
	Remark        string         `gorm:"type:varchar(255)" json:"remark"`
	CurentStepSeq int            `gorm:"type:int;not null" json:"curent_step_seq"`
	CreateBy      string         `gorm:"type:varchar(100)" json:"create_by"`
	CreateDtm     time.Time      `gorm:"autoCreateTime;<-:create" json:"create_date"`
	UpdateBy      string         `gorm:"type:varchar(100)" json:"update_by"`
	UpdateDTM     time.Time      `gorm:"autoUpdateTime;<-" json:"update_date"`
	ApprovalItem  []ApprovalItem `gorm:"foreignKey:ApprovalID;references:ID" json:"approval_item"`
}

func (Approval) TableName() string {
	return "approval"
}

type ApprovalItem struct {
	ID                     uuid.UUID                `gorm:"type:uuid;primaryKey;not null" json:"id"`
	ApprovalID             uuid.UUID                `gorm:"type:uuid;not null" json:"approval_id"`
	StepSeq                int                      `gorm:"type:int;not null" json:"step_seq"`
	IsCondition            bool                     `gorm:"type:boolean;not null" json:"is_condition"`
	Condition              string                   `gorm:"type:varchar(255)" json:"condition"`
	Status                 string                   `gorm:"type:varchar(20);not null" json:"status"`
	ActionBy               string                   `gorm:"type:varchar(100)" json:"action_by"`
	ActionDate             time.Time                `gorm:"type:timestamp" json:"action_date"`
	CreateBy               string                   `gorm:"type:varchar(100)" json:"create_by"`
	CreateDtm              time.Time                `gorm:"autoCreateTime;<-:create" json:"create_date"`
	UpdateBy               string                   `gorm:"type:varchar(100)" json:"update_by"`
	UpdateDTM              time.Time                `gorm:"autoUpdateTime;<-" json:"update_date"`
	ApprovalItemPermission []ApprovalItemPermission `gorm:"foreignKey:ApprovalItemID;references:ID" json:"approval_item_permission"`
}

func (ApprovalItem) TableName() string {
	return "approval_item"
}

type ApprovalItemPermission struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"id"`
	ApprovalItemID uuid.UUID `gorm:"type:uuid;not null" json:"approval_item_id"`
	UserCode       string    `gorm:"type:varchar(100)" json:"user_code"`
}

func (ApprovalItemPermission) TableName() string {
	return "approval_item_permission"
}

/* type Approval struct {
	ID             uuid.UUID  `json:"id"`
	ApproveCode    string     `json:"approve_code"`
	ApproveTopic   string     `json:"approve_topic"`
	DocumentType   string     `json:"document_type"`
	DocumentCode   string     `json:"document_code"`
	DocumentData   string     `json:"document_data"`
	ActionDate     *time.Time `json:"action_date"`
	Status         string     `json:"status"`
	Remark         string     `json:"remark"`
	CurrentStepSeq int        `json:"curent_step_seq"`
	CreateDate     *time.Time `json:"create_date"`
	CreateBy       string     `json:"create_by"`
	UpdateDate     *time.Time `json:"update_date"`
	UpdateBy       string     `json:"update_by"`
}

func (Approval) TableName() string { return "approval" }

type ApprovalItem struct {
	ID          uuid.UUID  `json:"id"`
	ApprovalID  uuid.UUID  `json:"approval_id"`
	StepSeq     int        `json:"step_seq"`
	IsCondition bool       `json:"is_condition"`
	Condition   string     `json:"condition"`
	Status      string     `json:"status"`
	ActionBy    string     `json:"action_by"`
	ActionDate  *time.Time `json:"action_date"`
	CreateDate  *time.Time `json:"create_date"`
	CreateBy    string     `json:"create_by"`
	UpdateDate  *time.Time `json:"update_date"`
	UpdateBy    string     `json:"update_by"`
}

func (ApprovalItem) TableName() string { return "approval_item" }

type ApprovalItemPermission struct {
	ID             uuid.UUID `json:"id"`
	ApprovalItemID uuid.UUID `json:"approval_item_id"`
	UserCode       string    `json:"user_code"`
}

func (ApprovalItemPermission) TableName() string { return "approval_item_permission" } */
