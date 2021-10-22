package subs

import (
	"errors"
	"net/http"

	ctx "github.com/yuriygr/go-posledstvie/context"

	"github.com/go-chi/render"
	"github.com/yuriygr/go-posledstvie/models"
	"github.com/yuriygr/go-posledstvie/utils"
)

// GetSubsite -
func (rs *Resource) GetSubsite(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.SubsiteRequest{
		CurrentUserID: session.ID,
	}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	subsite, err := rs.repository.GetSubsite(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}
	if err := render.Render(w, r, subsite); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

// Subscribe -
func (rs *Resource) Subscribe(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.SubsiteSubscribe{From: session.ID}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	// Check, exsist subsite or not
	if ok := rs.repository.ExistSubsite(request.To); !ok {
		render.Render(w, r, utils.ErrBadRequest(errors.New("Подсайт не существует")))
		return
	}

	if err := rs.repository.Subscribe(request); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Успешно подписался",
		AppCode:        18501,
	})
}

// Unsubscribe -
func (rs *Resource) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.SubsiteUnsubscribe{From: session.ID}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	// Check, exsist subsite or not
	if ok := rs.repository.ExistSubsite(request.To); !ok {
		render.Render(w, r, utils.ErrBadRequest(errors.New("Подсайт не существует")))
		return
	}

	if err := rs.repository.Unubscribe(request); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Успешно отписался",
		AppCode:        18500,
	})
}

// Mute -
func (rs *Resource) Mute(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Добавлен в черный список",
		AppCode:        18401,
	})
}

// Unmute -
func (rs *Resource) Unmute(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Удален из черного списка",
		AppCode:        18400,
	})
}

// Notification -
func (rs *Resource) Notification(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Тут поидее список подписчиков",
		AppCode:        18300,
	})
}

// Subscriptions -
func (rs *Resource) Subscriptions(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Тут поидее список подписок",
		AppCode:        18301,
	})
}

// Subscribers -
func (rs *Resource) Subscribers(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Тут поидее список подписчиков",
		AppCode:        18300,
	})
}

// Entries -
func (rs *Resource) Entries(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	request := &models.EntriesRequest{
		CurrentUserID: session.ID,
		PinnedOnTop:   true,

		Sort:  "e.created",
		Order: "desc",
	}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	entries, err := rs.repository.GetSubsiteEntries(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	if err := render.RenderList(w, r, models.NewEntryResponse(entries)); err != nil {
		render.Render(w, r, utils.ErrRender(err))
		return
	}
}

// Comments -
func (rs *Resource) Comments(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     "Тут поидее список подписчиков",
		AppCode:        18300,
	})
}
