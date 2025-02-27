package views

import (
	"net/http"

	"github.com/a-h/templ"
)

type CartProp struct {
	UserID    uint32
	ProductID uint32
	Size      string
	Quantity  uint32
	Name      string
	Price     uint32
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(cart([]CartProp{})).ServeHTTP(w, r)
}
