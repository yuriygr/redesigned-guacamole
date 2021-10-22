package user

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
	return &Resource{"/user", container, repository}
}

// Path - Путь
func (rs *Resource) Path() string {
	return rs.path
}

// Routes - Основные методы
func (rs *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.GetUser)
	})

	return r
}

// -- Методы

// GetUser -
func (rs *Resource) GetUser(w http.ResponseWriter, r *http.Request) {
	request := &models.SubsiteRequest{}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	user, err := rs.repository.GetUser(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}
	if err := render.Render(w, r, user); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

// Subscribe -
func (rs *Resource) Subscribe(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		StatusText: "Успешно подписался",
		AppCode:    18501,
	})
}

// Unsubscribe -
func (rs *Resource) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		StatusText: "Успешно отписался",
		AppCode:    18500,
	})
}

// Mute -
func (rs *Resource) Mute(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		StatusText: "Добавлен в черный список",
		AppCode:    18401,
	})
}

// Unmute -
func (rs *Resource) Unmute(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		StatusText: "Удален из черного списка",
		AppCode:    18400,
	})
}

// Subscriptions -
func (rs *Resource) Subscriptions(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		StatusText: "Тут поидее список подписок",
		AppCode:    18301,
	})
}

// Subscribers -
func (rs *Resource) Subscribers(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		StatusText: "Тут поидее список подписчиков",
		AppCode:    18300,
	})
}
