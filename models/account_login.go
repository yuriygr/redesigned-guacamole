package models

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/yuriygr/go-posledstvie/utils"
)

// AccountLogin - Чисто для логов
type AccountLogin struct {
	AccountID uint32 `db:"al.account_id"`
	IP        string `db:"al.ip"`
	Useragent string `db:"al.useragent"`
	Meta      struct{}
	Date      string `db:"al.date"`
}

// Bind - Bind HTTP request data and validate it
func (cul *AccountLogin) Bind(r *http.Request) error {
	cul.IP = ReadUserIP(r)
	cul.Useragent = r.UserAgent()
	cul.Date = time.Now().Format("2006-01-02 15:04:03")
	return nil
}

// Render - Render, wtf
func (ca *AccountLogin) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewLoginResponse - Вывод товаров годным списком
func NewLoginResponse(logins []*AccountLogin) []render.Renderer {
	list := []render.Renderer{}
	for _, login := range logins {
		list = append(list, login)
	}
	return list
}

// LoginRequest - Запрос списка авторизаций
type LoginRequest struct {
	Page, Limit uint32

	AccountID uint32
}

// Bind - Bind HTTP request data and validate it
func (ar *LoginRequest) Bind(r *http.Request) error {
	if page := r.URL.Query().Get("page"); page != "" {
		ar.Page = utils.Uint32(page)
	}

	if limit := r.URL.Query().Get("limit"); limit != "" {
		ar.Limit = utils.Uint32(limit)
	}

	return nil
}

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
