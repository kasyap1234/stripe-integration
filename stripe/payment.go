package stripe

import (
	"context"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/customer"
)

type CreateCustomerParameterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerResultResponse struct {
	Customer *stripe.Customer `json:"customer"`
}

//encore:api public method=POST path=/stripe/customers
func (s *Service) CreateCustomer(ctx context.Context, req CreateCustomerParameterRequest) (*CustomerResultResponse, error) {
	params := &stripe.CustomerParams{
		Email: stripe.String(req.Email),
		Name:  stripe.String(req.Name),
	}

	cust, err := s.stripeClient.Customers.New(params)
	if err != nil {
		return nil, err
	}

	return &CustomerResultResponse{
		Customer: cust,
	}, nil
}