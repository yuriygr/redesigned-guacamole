package models

import (
	"net/http"

	"github.com/go-chi/render"
)

// FastSearch -
type FastSearch struct {
	Value string `json:"value"`
	Type  string `json:"type"`
	UUID  string `json:"uuid"`
	URL   string `json:"url"`

	Payload struct {
	} `json:"payload"`
}

// Render - Render, wtf
func (ca *FastSearch) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewFastSearchResponse -
func NewFastSearchResponse(items []*FastSearch) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range items {
		list = append(list, item)
	}
	return list
}

// SearchRequest -
type SearchRequest struct {
	Query       string
	Order       string
	Page, Limit uint32
}

// Bind - Bind HTTP request data and validate it
func (pr *SearchRequest) Bind(r *http.Request) error {
	if query := r.URL.Query().Get("query"); query != "" {
		pr.Query = string(query)
	}

	return nil
}
