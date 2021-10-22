package widget

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
	return &Resource{"/widget", container, repository}
}

// Path - Путь
func (rs *Resource) Path() string {
	return rs.path
}

// Routes - Основные методы
func (rs *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/news", func(r chi.Router) {
		r.Get("/", rs.GetNews)
		r.Get("/more", rs.MoreNews)
	})

	return r
}

// -- Методы

// GetNews -
func (rs *Resource) GetNews(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, `{
		"items": [
			{ "title": "Утечка: Планируется к релизу новая версия Angular JS", "date":"1632823004", "url": "" },
			{ "title": "Вышла версия VueJS 3", "date":"1632820004", "url": "" },
			{ "title": "Ха наебал виджет новостей не работает", "date":"1632813004", "url": "" },
			{ "title": "И нет чинить не буду", "date":"1632823004", "url": "" },
			{ "title": "Ха напиздел еще раз, конечно починю", "date":"1632823004", "url": "" }
		]
	}`)
}

// MoreNews -
func (rs *Resource) MoreNews(w http.ResponseWriter, r *http.Request) {
	render.PlainText(w, r, `{
		"items": [
			{ "title": "Опа-на, еще что-то подгрузилось", "date":"1602823004", "url": "" }
		]
	}`)
}
