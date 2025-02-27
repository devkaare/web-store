package views

import (
	"net/http"

	"github.com/a-h/templ"
)

type cartProp struct {
	UserID    uint32
	ProductID uint32
	Size      string
	Quantity  uint32
	Name      string
	Price     uint32
	ImagePath string
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(cart([]cartProp{})).ServeHTTP(w, r)
}
