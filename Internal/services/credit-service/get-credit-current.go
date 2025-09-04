package creditService

import (
	"encoding/json"
	"errors"
	"fmt"
	"prime-erp-core/internal/db"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type GetCreditRequest struct {
	CustomerCodes []string `json:"customer_codes"`
}

type GetCreditResponse struct {
	CreditCustomers []CreditCustomer `json:"credit_customers"`
}

type CreditCustomer struct {
	CustomerCode string  `json:"customer_code"`
	IsActive     bool    `json:"is_active"`
	Credit       float64 `json:"credit"`
	Extra        float64 `json:"extra"`
	Used         float64 `json:"used"`
	Balance      float64 `json:"balance"`
}

func GetCreditCurrentAPI(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	req := GetCreditRequest{}

	if err := json.Unmarshal([]byte(jsonPayload), &req); err != nil {
		return nil, errors.New("failed to unmarshal JSON into struct: " + err.Error())
	}

	sqlx, err := db.ConnectSqlx(`prime_erp`)
	if err != nil {
		return nil, err
	}
	defer sqlx.Close()

	if len(req.CustomerCodes) == 0 {
		return nil, fmt.Errorf("require at least one customer code")
	}

	return GetCreditCurrent(sqlx, req)
}

func GetCreditCurrent(sqlx *sqlx.DB, req GetCreditRequest) (*GetCreditResponse, error) {
	res := GetCreditResponse{}
	customerCheck := map[string]bool{}
	customerStrs := []string{}

	for _, customer := range req.CustomerCodes {
		if !customerCheck[customer] {
			res.CreditCustomers = append(res.CreditCustomers, CreditCustomer{
				CustomerCode: customer,
				IsActive:     false,
				Credit:       0,
				Extra:        0,
				Used:         0,
				Balance:      0,
			})

			customerStrs = append(customerStrs, customer)
			customerCheck[customer] = true
		}
	}

	res, err := getCreditByCustomer(sqlx, res, customerStrs)
	if err != nil {
		return nil, err
	}

	res, err = getUsedByCustomer(sqlx, res, customerStrs)
	if err != nil {
		return nil, err
	}

	for i, customer := range res.CreditCustomers {
		customer.Balance = customer.Credit + customer.Extra - customer.Used
		res.CreditCustomers[i] = customer
	}

	return &res, nil
}

func getCreditByCustomer(sqlx *sqlx.DB, res GetCreditResponse, customerStrs []string) (GetCreditResponse, error) {
	if len(customerStrs) == 0 {
		return res, fmt.Errorf("no customer to check credit")
	}

	query := fmt.Sprintf(`
		select c.customer_code , coalesce(c.amount,0) credit_amount, coalesce (c.is_active,false) credit_is_active
			, coalesce(ce.amount, 0) extra_amount
		from credit c 
		left join credit_extra ce ON c.id = ce.credit_id  and ce.expire_dtm >= now() and ce.effective_dtm <= now()
		where 1=1
		and customer_code in ('%s')
	`, strings.Join(customerStrs, `','`))
	rows, err := db.ExecuteQuery(sqlx, query)
	if err != nil {
		return res, err
	}

	for _, row := range rows {
		customerCode := row["customer_code"].(string)
		creditAmount := row["credit_amount"].(float64)
		creditIsActive := row["credit_is_active"].(bool)
		extraAmount := row["extra_amount"].(float64)

		for i, customer := range res.CreditCustomers {
			if customer.CustomerCode == customerCode {
				customer.IsActive = creditIsActive
				customer.Credit = creditAmount
				customer.Extra = extraAmount

				res.CreditCustomers[i] = customer
			}
		}
	}

	return res, nil
}

func getUsedByCustomer(sqlx *sqlx.DB, res GetCreditResponse, customerStrs []string) (GetCreditResponse, error) {
	if len(customerStrs) == 0 {
		return res, fmt.Errorf("no customer to check used credit")
	}

	query := fmt.Sprintf(`
		select customer_code, sum(total) used_amount
		from invoice
		where 1=1
		and customer_code in ('%s')
		and is_canceled = false
		and is_approved = true
		and invoice_dtm >= now() - interval '1 year'
		group by customer_code
	`, strings.Join(customerStrs, `','`))
	rows, err := db.ExecuteQuery(sqlx, query)
	if err != nil {
		return res, err
	}

	for _, row := range rows {
		customerCode := row["customer_code"].(string)
		usedAmount := row["used_amount"].(float64)
		for i, customer := range res.CreditCustomers {
			if customer.CustomerCode == customerCode {
				customer.Used = usedAmount
				res.CreditCustomers[i] = customer
			}
		}
	}

	return res, nil
}
