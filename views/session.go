package views

import (
	"net/http"

	"github.com/a-h/templ"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(signUp()).ServeHTTP(w, r)
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(signIn()).ServeHTTP(w, r)
}
