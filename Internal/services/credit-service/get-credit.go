package creditService

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

type GetCreditRequest struct {
	CustomerCodes []string `json:"customer_codes"`
}

type GetCreditResponse struct {
	CreditCustomers []CreditCustomer `json:"credit_customers"`
}

type CreditCustomer struct {
	CustomerCode  string  `json:"customer_code"`
	CreditLimit   float64 `json:"credit_limit"`
	UsedCredit    float64 `json:"used_credit"`
	BalanceCredit float64 `json:"balance_credit"`
}

func GetCredit(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	req := GetCreditRequest{}

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	return GetCreditLogic(req)
}

func GetCreditLogic(req GetCreditRequest) (*GetCreditResponse, error) {
	res := GetCreditResponse{}

	return &res, nil
}
