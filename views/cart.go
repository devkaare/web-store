package views

import (
	"net/http"

	"github.com/a-h/templ"
)

type cartProp struct {
	UserID    int
	ProductID int
	Size      string
	Quantity  int
	Name      string
	Price     int
	ImagePath string
}

func CartHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(cart([]cartProp{})).ServeHTTP(w, r)
}
