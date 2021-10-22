package timeline

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
	return &Resource{"/timeline", container, repository}
}

// Path - Путь
func (rs *Resource) Path() string {
	return rs.path
}

// Routes - Основные методы
func (rs *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.Entries)

	return r
}

// -- Методы

// Entries -
func (rs *Resource) Entries(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.EntriesRequest{
		CurrentUserID: session.ID,
		PinnedOnTop:   false,

		Sort:  "e.created",
		Order: "desc",
	}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	entries, err := rs.repository.GetTimelineEntries(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	if err := render.RenderList(w, r, models.NewEntryResponse(entries)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}
