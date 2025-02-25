package views

import (
	"net/http"

	"github.com/a-h/templ"
)

func CartHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(cart([]CartProp{})).ServeHTTP(w, r)
}
