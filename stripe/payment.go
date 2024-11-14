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
type SourceTypes struct {
    Card int64 `json:"card"`
}

type BalanceAmount struct {
    Amount      int64       `json:"amount"`
    Currency    string      `json:"currency"`
    SourceTypes SourceTypes `json:"source_types,omitempty"`
}

type BalanceResult struct {
    Object          string          `json:"object"`
    Available       []BalanceAmount `json:"available"`
    ConnectReserved []BalanceAmount `json:"connect_reserved"`
    Livemode        bool           `json:"livemode"`
    Pending         []BalanceAmount `json:"pending"`
}


//encore:api public method=GET path=/v1/balance
func (s *Service) GetBalance(ctx context.Context) (*BalanceResult, error) {
    params := &stripe.BalanceParams{}
    
    result, err := s.stripeClient.Balance.Get(params)
    if err != nil {
        return nil, err
    }

    // Convert Available amounts
    available := make([]BalanceAmount, len(result.Available))
    for i, amt := range result.Available {
        available[i] = BalanceAmount{
            Amount:   amt.Amount,
            Currency: string(amt.Currency),
            SourceTypes: SourceTypes{
                Card: amt.SourceTypes["card"],
            },
        }
    }

    // Convert ConnectReserved amounts
    connectReserved := make([]BalanceAmount, len(result.ConnectReserved))
    for i, amt := range result.ConnectReserved {
        connectReserved[i] = BalanceAmount{
            Amount:   amt.Amount,
            Currency: string(amt.Currency),
        }
    }

    // Convert Pending amounts
    pending := make([]BalanceAmount, len(result.Pending))
    for i, amt := range result.Pending {
        pending[i] = BalanceAmount{
            Amount:   amt.Amount,
            Currency: string(amt.Currency),
            SourceTypes: SourceTypes{
                Card: amt.SourceTypes["card"],
            },
        }
    }

    return &BalanceResult{
        Object:          result.Object,
        Available:       available,
        ConnectReserved: connectReserved,
        Livemode:       result.Livemode,
        Pending:        pending,
    }, nil
}



//encore:api public method=POST path=/v1/customers
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
