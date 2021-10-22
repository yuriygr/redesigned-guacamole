package context

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/yuriygr/go-posledstvie/services"
	"github.com/yuriygr/go-posledstvie/utils"
)

// AuthSessionKey - Key for context
type AuthSessionKey struct{}

// AuthSessionCtx - Контекст с авторизацией пользователя.
// В целом, у меня нет нареканий к данному коду.
func AuthSessionCtx(session *services.Session) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userSession, _ := session.AuthStart(w, r)
			ctx := context.WithValue(r.Context(), AuthSessionKey{}, userSession)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// AuthAccessCtx - Моя самая простая реализация защиты доступа к личным данным(?).
func AuthAccessCtx(session *services.Session) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userSession, _ := session.AuthStart(w, r)
			if !userSession.Auth {
				render.Render(w, r, utils.ErrUnauthorized())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
