package models

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/yuriygr/go-mlh/formatter"
	"github.com/yuriygr/go-posledstvie/utils"
)

// SubsRequest - Запрос
type SubsRequest struct {
	CurrentUserID uint32 `db:"r.current_user_id"`

	Sort, Order string
	Page, Limit uint32
	Type        string
}

// Bind - Bind HTTP request data and validate it
func (pr *SubsRequest) Bind(r *http.Request) error {

	if sort := r.URL.Query().Get("sort"); sort != "" {
		pr.Sort = formatter.EscapeString(sort)
	}

	if order := r.URL.Query().Get("order"); order != "" {
		pr.Order = formatter.EscapeString(order)
	}

	if page := r.URL.Query().Get("page"); page != "" {
		pr.Page = utils.Uint32(page)
	}

	if limit := r.URL.Query().Get("limit"); limit != "" {
		pr.Limit = utils.Uint32(limit)
	}

	return nil
}

// SubsiteRequest -
type SubsiteRequest struct {
	CurrentUserID uint32 `db:"r.current_user_id"`

	ID   uint32 `db:"r.subsite_id"`
	Slug string `db:"r.slug"`
}

// Render - Render, wtf
func (s *SubsiteRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind - Bind HTTP request data and validate it
func (pr *SubsiteRequest) Bind(r *http.Request) error {
	param := formatter.EscapeString(chi.URLParam(r, "param"))

	if _, err := strconv.Atoi(param); err == nil {
		pr.ID = utils.Uint32(param)
	} else {
		pr.Slug = param
	}

	return nil
}

// SubsiteSubscribe -
type SubsiteSubscribe struct {
	From uint32 `db:"r.from"`
	To   uint32 `db:"r.to"`
}

// Bind - Bind HTTP request data and validate it
func (pr *SubsiteSubscribe) Bind(r *http.Request) error {
	pr.To = utils.Uint32(formatter.EscapeString(chi.URLParam(r, "param")))

	return nil
}

// SubsiteUnsubscribe -
type SubsiteUnsubscribe struct {
	From uint32 `db:"r.from"`
	To   uint32 `db:"r.to"`
}

// Bind - Bind HTTP request data and validate it
func (pr *SubsiteUnsubscribe) Bind(r *http.Request) error {
	pr.To = utils.Uint32(formatter.EscapeString(chi.URLParam(r, "param")))

	return nil
}
