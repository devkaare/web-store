package handler

import (
	"context"
	"log"
	"net/http"
	// "github.com/devkaare/web-store/model"
	// "github.com/devkaare/web-store/repository/shopping_session"
)

func (a *Authentication) SetUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if checkIfErrNoCookie(err) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		sessionID := cookie.Value

		session, err := sessionHandler.Repo.GetSessionBySessionID(sessionID)
		if checkIfErrNoRows(err) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if isExpired(session) {
			if err := sessionHandler.Repo.DeleteSessionBySessionID(sessionID); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", session.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// func (a *Authentication) CreateShoppingSessionID(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		shoppingSession := &model.ShoppingSession{}
//
// 		shoppingSessionID, err := shoppingSessionHandler.Repo.CreateShoppingSession(shoppingSession)
// 		check(err)
//
// 		ctx := context.WithValue(r.Context(), "shopping_session_id", shoppingSessionID)
//
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
