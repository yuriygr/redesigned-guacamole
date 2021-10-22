package search

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yuriygr/go-posledstvie/container"
	"github.com/yuriygr/go-posledstvie/models"
	"github.com/yuriygr/go-posledstvie/repository"
	"github.com/yuriygr/go-posledstvie/utils"
)

// Resource - Ресурс каталога
type Resource struct {
	path       string
	container  *container.Container
	repository *repository.Repository
}

// BuildResource - Создаем новый экземпляр ресурса
func BuildResource(repository *repository.Repository, container *container.Container) *Resource {
	return &Resource{"/search", container, repository}
}

// Path - Путь
func (rs *Resource) Path() string {
	return rs.path
}

// Routes - Основные методы
func (rs *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/fast", rs.FastSearch)

	return r
}

// -- Методы

// FastSearch -
func (rs *Resource) FastSearch(w http.ResponseWriter, r *http.Request) {
	request := &models.SearchRequest{
		Order: "desc",
		Page:  1,
		Limit: 5,
	}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	results, err := rs.repository.GetFastSearch(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	if err := render.RenderList(w, r, models.NewFastSearchResponse(results)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}
