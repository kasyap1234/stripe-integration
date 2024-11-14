package stripe

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/client"
)

//encore:service
type Service struct {
	stripeClient *client.API
}

// initService initializes the Stripe service
var secrets struct {
	StripeKey string
}

func initService() (*Service, error) {
	stripe.Key = secrets.StripeKey
	stripeClient := &client.API{}
	stripeClient.Init(stripe.Key, nil)
	return &Service{stripeClient: stripeClient}, nil
}
