package priceService

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

type GetPaymentTermRequest struct {
}

func GetPaymentTerm(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	req := GetPaymentTermRequest{}

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	return nil, nil
}
