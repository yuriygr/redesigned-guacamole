package context

import (
	"net/http"

	"github.com/yuriygr/go-posledstvie/container"
)

// TODO: Вынести отдельно КЛЮЧИ чтобы можно было использовать в ресурсах

// APIAccess - Моя самая простая реализация защиты доступа к API.
// Конечно, существуют готовые решения данных проблем,
// но я все же учусь. Так ведь?
func APIAccess(container *container.Container) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
