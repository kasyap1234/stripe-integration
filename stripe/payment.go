package stripe

import (
    "context"
    "fmt"  
    "github.com/stripe/stripe-go/v81"
)

type CreateCustomerParameterRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CustomerResultResponse struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Created int64  `json:"created"`
}

//encore:api public method=POST path=/stripe/customers
func (s *Service) CreateCustomer(ctx context.Context, req CreateCustomerParameterRequest) (*CustomerResultResponse, error) {
    if s.stripeClient == nil {
        return nil, fmt.Errorf("stripe client not initialized")
    }

    params := &stripe.CustomerParams{
        Email: stripe.String(req.Email),
        Name:  stripe.String(req.Name),
    }

    cust, err := s.stripeClient.Customers.New(params)
    if err != nil {
        return nil, fmt.Errorf("failed to create customer: %w", err)
    }

    return &CustomerResultResponse{
        ID:      cust.ID,
        Name:    cust.Name,
        Email:   cust.Email,
        Created: cust.Created,
    }, nil
}
