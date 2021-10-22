package rating

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yuriygr/go-posledstvie/container"
	"github.com/yuriygr/go-posledstvie/repository"
)

// Resource - Ресурс каталога
type Resource struct {
	path       string
	container  *container.Container
	repository *repository.Repository
}

// BuildResource - Создаем новый экземпляр ресурса
func BuildResource(repository *repository.Repository, container *container.Container) *Resource {
	return &Resource{"/rating", container, repository}
}

// Path - Путь
func (rs *Resource) Path() string {
	return rs.path
}

// Routes - Основные методы
func (rs *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.GetSubs)
	r.Get("/3month", rs.GetSubs)
	r.Get("/all", rs.GetSubs)

	return r
}

// -- Методы

// GetSubs - Подписки
func (rs *Resource) GetSubs(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, `НЕ ну а х=ули`)
}
