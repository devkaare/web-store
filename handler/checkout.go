package handler

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
)

type Checkout struct{}

func NewCheckoutHandler() *Checkout {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	return &Checkout{}
}

func (c *Checkout) CreateCheckoutSession(w http.ResponseWriter, r *http.Request) {
	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("usd"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Service"),
					},
					UnitAmount: stripe.Int64(3000),
				},
				Quantity: stripe.Int64(1),
			},
		},
		SuccessURL: stripe.String("http://localhost:3000/checkout/success"),
		CancelURL:  stripe.String("http://localhost:3000/checkout/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, s.URL, http.StatusSeeOther)
}
