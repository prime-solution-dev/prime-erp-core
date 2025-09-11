package approvalService

import (
	"encoding/json"
	"errors"
	models "prime-erp-core/internal/models"
	repositoryApproval "prime-erp-core/internal/repositories/approval"
	authenticationService "prime-erp-core/internal/services/authentication-service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type JSONB struct {
	json.RawMessage
}

func CreateApproval(ctx *gin.Context, jsonPayload string) (interface{}, error) {

	var req []models.Approval

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	approvalValue := []models.Approval{}
	approvalItemValue := []models.ApprovalItem{}
	approvalItemPermissionValue := []models.ApprovalItemPermission{}
	approvalIDForReturn := []uuid.UUID{}
	mdiItemCode := []string{}

	for _, approval := range req {
		mdiItemCode = append(mdiItemCode, approval.MDItemCode)
	}

	requestData := map[string]interface{}{
		"md_item_code": mdiItemCode,
		"action_code":  []string{"APPROVE"},
	}
	requester, errGetRequester := authenticationService.GetRequester(requestData)
	if errGetRequester != nil {
		return nil, errGetRequester
	}

	for i, approval := range req {

		approvalID := uuid.New()
		req[i].ID = approvalID
		approvalIDForReturn = append(approvalIDForReturn, approvalID)

		for o := range approval.ApprovalItem {
			approvalItemID := uuid.New()
			req[i].ApprovalItem[o].ID = approvalItemID
			req[i].ApprovalItem[o].ApprovalID = approvalID
			req[i].ApprovalItem[o].StepSeq = 1
			req[i].ApprovalItem[o].IsCondition = false
			jsonDataApprovalItem, _ := json.Marshal(req[i].ApprovalItem[o])
			req[i].ApprovalItem[o].Condition = jsonDataApprovalItem
			req[i].ApprovalItem[o].Status = "PENDING"
			req[i].ApprovalItem[o].ActionDate = time.Now()

			for _, requesterValue := range requester {
				approvalItemPermissionID := uuid.New()
				newApprovalItemPermission := models.ApprovalItemPermission{
					ID:             approvalItemPermissionID,
					ApprovalItemID: approvalItemID,
					UserCode:       requesterValue.RequesterID,
				}
				approvalItemPermissionValue = append(approvalItemPermissionValue, newApprovalItemPermission)
			}
			req[i].ApprovalItem[o].ApprovalItemPermission = []models.ApprovalItemPermission{}
			approvalItemValue = append(approvalItemValue, req[i].ApprovalItem[o])

		}

		if req[i].ApproveCode != "" {
			req[i].ApproveCode = approval.ApproveCode
		} else {
			req[i].ApproveCode = uuid.New().String()
		}
		jsonDataApproval, _ := json.Marshal(req[i])
		req[i].DocumentData = jsonDataApproval
		req[i].ApprovalItem = []models.ApprovalItem{}

		approvalValue = append(approvalValue, req[i])
	}

	errCreateApproval := repositoryApproval.CreateApproval(approvalValue, approvalItemValue, approvalItemPermissionValue)
	if errCreateApproval != nil {
		return nil, errCreateApproval
	}

	return map[string]interface{}{
		"id":      approvalIDForReturn,
		"status":  "success",
		"message": "Approval create successfully",
	}, nil
}
