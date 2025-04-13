package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/devkaare/web-store/model"
	// "github.com/google/uuid"
)

// func (a *Authentication) SessionMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("session_token")
// 		if err != nil {
// 			if err == http.ErrNoCookie {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				return
// 			}
// 			log.Println(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
//
// 		sessionID := cookie.Value
//
// 		session, err := sessionHandler.Repo.GetSessionBySessionID(sessionID)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				w.WriteHeader(http.StatusUnauthorized)
// 				return
// 			}
// 			log.Println(err)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
//
// 		ctx := context.WithValue(r.Context(), "user_id", session.UserID)
//
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func (a *Authentication) ShoppingSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var session *model.Session
		var shoppingSession *model.ShoppingSession

		cookie, err := r.Cookie("session_token")
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		session, err = a.SessionRepo.GetSessionBySessionID(cookie.Value)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		shoppingSession, err = a.ShoppingSessionRepo.GetShoppingSessionBySessionID(session.SessionID)
		if err == sql.ErrNoRows {
			shoppingSession = &model.ShoppingSession{SessionID: session.SessionID, Total: 0}

			shoppingSessionID, err := a.ShoppingSessionRepo.CreateShoppingSession(shoppingSession)
			if err != nil {
				log.Printf("ShoppingSessionMiddleware: error creating shopping session: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			shoppingSession.ShoppingSessionID = shoppingSessionID
		} else if err != nil {
			log.Printf("ShoppingSessionMiddleware: error fetching shopping session by shopping session ID: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("Set shopping session id: %d\n", shoppingSession.ShoppingSessionID)
		ctx := context.WithValue(r.Context(), "shopping_session_id", shoppingSession.ShoppingSessionID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
