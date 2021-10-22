package auth

import (
	"errors"
	"net/http"

	"github.com/yuriygr/go-mlh/formatter"
	ctx "github.com/yuriygr/go-posledstvie/context"
	"github.com/yuriygr/go-posledstvie/models"
	"github.com/yuriygr/go-posledstvie/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yuriygr/go-posledstvie/container"
	"github.com/yuriygr/go-posledstvie/utils"
)

const (
	errUserNotFound    = "Invalid username or password"
	errAlreadyLogged   = "You are already logged in, where will you go again?"
	errUserIsBanned    = "Sorry Mario, the Princess is in another castle"
	errUserIsDeleted   = "This account does not exist"
	errUserIsNotActive = "This account does not active"
	succLogin          = "You are successfully logged in."
	succLogout         = "You are successfully log out."
)

// Resource - Ресурс оформления заказа
type Resource struct {
	path       string
	container  *container.Container
	repository *repository.Repository
}

// BuildResource - Создаем новый экземпляр ресурса
func BuildResource(repository *repository.Repository, container *container.Container) *Resource {
	return &Resource{"/auth", container, repository}
}

// Path - Путь
func (rs *Resource) Path() string {
	return rs.path
}

// Routes - Основные методы
func (rs *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/session", rs.Session)
	r.Post("/login", rs.Login)
	r.Post("/register", rs.Register)
	r.Post("/forgot", rs.Forgot)
	r.Post("/logout", rs.Logout)

	return r
}

// -- Методы

// Session - Сесси
func (rs *Resource) Session(w http.ResponseWriter, r *http.Request) {
	session, _ := rs.container.Session.AuthGet(w, r)
	render.Render(w, r, session)
}

// Login - Авторизация
func (rs *Resource) Login(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	if session.Auth {
		err := errors.New(errAlreadyLogged)
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	email := r.FormValue("email")
	email = formatter.EscapeString(email)
	account, err := rs.repository.GetAccountByEmail(email)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(errors.New(errUserNotFound)))
		return
	}

	password := r.FormValue("password")
	if !utils.CheckPasswordHash(password, account.Password) {
		render.Render(w, r, utils.ErrBadRequest(errors.New(errUserNotFound)))
		return
	}

	// Мы не шпионим, просто секюрности нада насяльника
	account_login := &models.AccountLogin{AccountID: account.ID}
	if err := account_login.Bind(r); err != nil {
		rs.container.Logger.Error(err)
	}
	err = rs.repository.CreateAccountLogin(account_login)
	if err != nil {
		rs.container.Logger.Error(err)
	}

	// Получаем _текущий_ подсайт пользователя для передачи в сессию
	subsite, err := rs.repository.GetUser(&models.SubsiteRequest{
		ID: account.ID,
	})
	if err != nil {
		rs.container.Logger.Error(err)
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	// Ok, user exist and password correct...
	// Let's create session!

	session, err = rs.container.Session.AuthCreate(w, r, subsite)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     succLogin,
		Payload:        session,
	})
}

// Register - Регистрация
func (rs *Resource) Register(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(ctx.AuthSessionKey{}).(*models.SessionResponse)
	if session.Auth {
		err := errors.New(errAlreadyLogged)
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	request := &models.Account{}
	if err := request.Bind(r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	account, err := rs.repository.CreateAccount(request)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	// Мы не шпионим, просто секюрности нада насяльника
	account_login := &models.AccountLogin{AccountID: account.ID}
	if err := account_login.Bind(r); err != nil {
		rs.container.Logger.Error(err)
	}
	err = rs.repository.CreateAccountLogin(account_login)
	if err != nil {
		rs.container.Logger.Error(err)
	}

	// Ok, user exist and password correct...
	// Let's create session!

	// Получаем _текущий_ подсайт пользователя для передачи в сессию
	subsite, err := rs.repository.GetUser(&models.SubsiteRequest{
		ID: account.ID,
	})
	if err != nil {
		rs.container.Logger.Error(err)
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	session, err = rs.container.Session.AuthCreate(w, r, subsite)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     succLogin,
		Payload:        session,
	})
}

// Forgot - Забыли пароль
func (rs *Resource) Forgot(w http.ResponseWriter, r *http.Request) {}

// Logout - Выход
func (rs *Resource) Logout(w http.ResponseWriter, r *http.Request) {
	if err := rs.container.Session.Logout(w, r); err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	// Let's create new session!

	session, err := rs.container.Session.AuthStart(w, r)
	if err != nil {
		render.Render(w, r, utils.ErrBadRequest(err))
		return
	}

	render.Render(w, r, &utils.SuccessResponse{
		HTTPStatusCode: 200,
		StatusText:     succLogout,
		Payload:        session,
	})
}
