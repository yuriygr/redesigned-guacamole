package main

import (
	"encoding/gob"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/yuriygr/go-posledstvie/container"
	ctx "github.com/yuriygr/go-posledstvie/context"
	"github.com/yuriygr/go-posledstvie/models"
	"github.com/yuriygr/go-posledstvie/repository"
	"github.com/yuriygr/go-posledstvie/services"

	"github.com/yuriygr/go-posledstvie/resources/auth"
	"github.com/yuriygr/go-posledstvie/resources/entry"
	"github.com/yuriygr/go-posledstvie/resources/rating"
	"github.com/yuriygr/go-posledstvie/resources/search"
	"github.com/yuriygr/go-posledstvie/resources/subs"
	"github.com/yuriygr/go-posledstvie/resources/timeline"
	"github.com/yuriygr/go-posledstvie/resources/user"
	"github.com/yuriygr/go-posledstvie/resources/widget"

	"github.com/yuriygr/go-posledstvie/utils"

	_ "github.com/go-sql-driver/mysql"
)

// APIVersion1 - Const for api versions
const APIVersion1 = "v1"

// NewRouter - Создаем новый роутер
func NewRouter(config *services.Config) *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   config.CORS.AllowedOrigins,
		AllowedMethods:   config.CORS.AllowedMethods,
		AllowedHeaders:   config.CORS.AllowedHeaders,
		ExposedHeaders:   config.CORS.ExposedHeaders,
		AllowCredentials: config.CORS.AllowCredentials,
		MaxAge:           config.CORS.MaxAge,
	})

	r.Use(cors.Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, utils.ErrMethodNotAllowed())
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, utils.ErrMethodNotAllowed())
	})

	return r
}

// Resource - Интерфейс ресурса
type Resource interface {
	Path() string
	Routes() chi.Router
}

// Да, все это ради красивого main(), как ты узнал?
// UDP: Господи, я люблю GO.
func main() {
	gob.Register(models.SessionResponse{})

	container := container.NewContainer()
	container.Config.Application.CurrentVersion = APIVersion1

	repository := repository.NewRepository(container)

	router := NewRouter(container.Config)

	router.Use(ctx.APIAccess(container))
	router.Use(ctx.AuthSessionCtx(container.Session))

	// Приветствуем мир
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		container.Logger.Info(container.Config.Application.FirstMessage)
		render.Render(w, r, utils.SuccessOK(container.Config.Application.FirstMessage))
	})

	router.Route("/{version}", func(r chi.Router) {
		// Нам необходимо как-то отслеживать текущую версию API
		// Чем контекст не иделаьное решение?
		r.Use(ctx.APIVersion("version"))

		// @todo Переписать на фабрику
		resources := []Resource{
			auth.BuildResource(repository, container),
			rating.BuildResource(repository, container),
			timeline.BuildResource(repository, container),
			entry.BuildResource(repository, container),
			subs.BuildResource(repository, container),
			user.BuildResource(repository, container),
			widget.BuildResource(repository, container),
			search.BuildResource(repository, container),
		}
		for _, resource := range resources {
			r.Mount(resource.Path(), resource.Routes())
		}
	})

	http.ListenAndServe(":"+container.Config.Application.Port, router)
}
