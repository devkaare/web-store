package views

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/devkaare/web-store/views/components"
)

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(components.Base()).ServeHTTP(w, r)
}
