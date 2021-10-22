package entry

import (
	"net/http"

	ctx "github.com/yuriygr/go-posledstvie/context"
	"github.com/yuriygr/go-posledstvie/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yuriygr/go-posledstvie/container"
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
	return &Resource{"/entry", container, repository}
}

// Path - Путь
func (rs *Resource) Path() string {
	return rs.path
}

// Routes - Основные методы
func (rs *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", rs.Entry)
		r.Get("/comments", rs.Comments)

		r.Route("/", func(r chi.Router) {
			r.Use(ctx.AuthAccessCtx(rs.container.Session))

			r.Post("/mute", rs.Mute)
			r.Post("/unmute", rs.Unmute)

			r.Post("/like", rs.Like)
			r.Post("/dislike", rs.Dislike)
		})
	})

	return r
}

// -- Методы

// Entry -
func (rs *Resource) Entry(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.EntryRequest{
		CurrentUserID: session.ID,
	}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	entry, err := rs.repository.GetEntry(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	render.Render(w, r, entry)
}

// Comments -
func (rs *Resource) Comments(w http.ResponseWriter, r *http.Request) {
	//session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	//request := &models.CommentsRequest{
	//	CurrentUserID: session.ID,
	//}
	//if err := request.Bind(r); err != nil {
	//	render.Render(w, r, utils.ErrBadRequest(err))
	//	return
	//}
	//
	//entries, err := rs.repository.GetEntries(request)
	//if err != nil {
	//	render.Render(w, r, utils.ErrBadRequest(err))
	//	return
	//}
	//
	//if err := render.RenderList(w, r, models.NewEntryResponse(entries)); err != nil {
	//	render.Render(w, r, utils.ErrRender(err))
	//	return
	//}
}

// Mute -
func (rs *Resource) Mute(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Скрыли (нет)",
		AppCode:        18300,
	})
}

// Unmute -
func (rs *Resource) Unmute(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Раскрыли (нет)",
		AppCode:        18300,
	})
}

// Like -
func (rs *Resource) Like(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.EntryLikeRequest{
		CurrentUserID: session.ID,
	}

	vote, err := rs.repository.EntryLike(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Успешно",
		AppCode:        18300,
		Payload:        vote,
	})
}

// Dislike -
func (rs *Resource) Dislike(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.EntryLikeRequest{
		CurrentUserID: session.ID,
	}

	vote, err := rs.repository.EntryDislike(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Успешно",
		AppCode:        18300,
		Payload:        vote,
	})
}
