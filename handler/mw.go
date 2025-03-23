package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/devkaare/web-store/model"
	"github.com/google/uuid"
)

func (a *Authentication) SessionMiddleware(next http.Handler) http.Handler {
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

func (a *Authentication) ShoppingSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var session *model.Session
		var shoppingSession *model.ShoppingSession

		cookie, err := r.Cookie("session_token")
		if err == nil {
			session, err = sessionHandler.Repo.GetSessionBySessionID(cookie.Value)
			if err == sql.ErrNoRows {
				session = nil
			} else if err != nil {
				log.Printf("ShoppingSessionMiddleware: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		if session == nil {
			sessionID := uuid.New().String()
			session := &model.Session{SessionID: sessionID}

			if err := sessionHandler.Repo.CreateSession(session); err != nil {
				log.Printf("ShoppingSessionMiddleware: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:  "session_token",
				Value: session.SessionID,
			})
		}

		shoppingSession, err = shoppingSessionHandler.Repo.GetShoppingSessionBySessionID(session.SessionID)
		if err != nil {
			log.Printf("ShoppingSessionMiddleware: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		shoppingSession = &model.ShoppingSession{SessionID: session.SessionID, Total: 0}

		shoppingSessionID, err := shoppingSessionHandler.Repo.CreateShoppingSession(shoppingSession)
		if err != nil {
			log.Printf("ShoppingSessionMiddleware: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		shoppingSession.ShoppingSessionID = shoppingSessionID

		ctx := context.WithValue(r.Context(), "shopping_session_id", shoppingSession.ShoppingSessionID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
