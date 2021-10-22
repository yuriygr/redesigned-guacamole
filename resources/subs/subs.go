package subs

import (
	"net/http"

	ctx "github.com/yuriygr/go-posledstvie/context"

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
	return &Resource{"/subsite", container, repository}
}

// Path - Путь
func (rs *Resource) Path() string {
	return rs.path
}

// Routes - Основные методы
func (rs *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.GetSubs)
	r.Get("/recommendations", rs.GetSubsRecommendations)
	r.Get("/companies", rs.GetSubsCompanies)

	r.Route("/{param}", func(r chi.Router) {
		r.Get("/", rs.GetSubsite)

		// Обязательная авторизация, йоу
		r.Route("/", func(r chi.Router) {
			r.Use(ctx.AuthAccessCtx(rs.container.Session))

			r.Post("/subscribe", rs.Subscribe)

			r.Post("/unsubscribe", rs.Unsubscribe)
			r.Post("/notifications", rs.Notification)

			r.Post("/mute", rs.Mute)
			r.Post("/unmute", rs.Unmute)
		})

		r.Get("/entries", rs.Entries)
		r.Get("/comments", rs.Comments)
		r.Get("/subscriptions", rs.Subscriptions)
		r.Get("/subscribers", rs.Subscribers)

	})

	return r
}

// -- Методы

// GetSubs - Подписки
func (rs *Resource) GetSubs(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.SubsRequest{ // Initial data
		CurrentUserID: session.ID,
		Sort:          "s.subsite_id",
		Order:         "desc",
		Page:          1,
		Limit:         100,
		Type:          "All",
	}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	subs, err := rs.repository.GetSubsites(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	if err := render.RenderList(w, r, models.NewMiniSubsResponse(subs)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

// GetSubsRecommendations -
func (rs *Resource) GetSubsRecommendations(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.SubsRequest{ // Initial data
		CurrentUserID: session.ID,
		Order:         "desc",
		Page:          1,
		Limit:         100,
		Type:          "Recommendations",
	}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	subs, err := rs.repository.GetSubsites(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	if err := render.RenderList(w, r, models.NewMiniSubsResponse(subs)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

// GetSubsCompanies -
func (rs *Resource) GetSubsCompanies(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.SubsRequest{ // Initial data
		CurrentUserID: session.ID,
		Order:         "desc",
		Page:          1,
		Limit:         100,
		Type:          "Companies",
	}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	subs, err := rs.repository.GetSubsites(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	if err := render.RenderList(w, r, models.NewMiniSubsResponse(subs)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}
