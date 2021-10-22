package context

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

// APIVersionKey - Key for context
type APIVersionKey struct{}

// APIVersion - Context API version
func APIVersion(param string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), APIVersionKey{}, chi.URLParam(r, param))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
