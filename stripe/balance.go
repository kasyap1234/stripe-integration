package stripe

import (
	"context"

	"github.com/stripe/stripe-go/v81"
)

type BalanceTransactionResponse struct {
	ID                string        `json:"id"`
	Object            string        `json:"object"`
	Amount            int64         `json:"amount"`
	AvailableOn       int64         `json:"available_on"`
	Created           int64         `json:"created"`
	Currency          string        `json:"currency"`
	Description       interface{}   `json:"description"`
	ExchangeRate      interface{}   `json:"exchange_rate"`
	Fee               int64         `json:"fee"`
	FeeDetails        []interface{} `json:"fee_details"`
	Net               int64         `json:"net"`
	ReportingCategory string        `json:"reporting_category"`
	Source            interface{}   `json:"source"`
	Status            string        `json:"status"`
	Type              string        `json:"type"`
}

//encore:api public method=GET  path=v1/balance_transactions/:id
func BalanceTransactions(context context.Context, s *Service, id string) (*BalanceTransactionResponse, error) {
	params := &stripe.BalanceTransactionParams{}
	result, err := s.stripeClient.BalanceTransactions.Get(id, params)
	if err != nil {
		return nil, err
	}

	feeDetails := make([]interface{}, len(result.FeeDetails))
	for i, fd := range result.FeeDetails {
		feeDetails[i] = fd
	}

	return &BalanceTransactionResponse{
		ID:                result.ID,
		Object:            result.Object,
		Amount:            result.Amount,
		AvailableOn:       result.AvailableOn,
		Created:           result.Created,
		Currency:          string(result.Currency),
		Description:       result.Description,
		ExchangeRate:      result.ExchangeRate,
		Fee:               result.Fee,
		FeeDetails:        feeDetails,
		Net:               result.Net,
		ReportingCategory: string(result.ReportingCategory),
		Source:            result.Source,
		Status:            string(result.Status),
		Type:              string(result.Type),
	}, nil
}
