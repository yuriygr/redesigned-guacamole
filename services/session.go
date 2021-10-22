package services

/**
 * В будущем надо будет описать, почему я сделал так, а не иначе
 */

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/yuriygr/go-loggy"
	"github.com/yuriygr/go-mlh/formatter"
	"github.com/yuriygr/go-posledstvie/models"
	"gopkg.in/boj/redistore.v1"
)

var (
	onceSession sync.Once
	session     *Session
)

// Default keys
var (
	authSessionID = "sid"
)

// NewSession - Создаем экземпляр хранилища
func NewSession(config *Config, logger *loggy.Logger) *Session {
	authSessionID = config.Session.Key + "_sid"

	onceSession.Do(func() {
		address := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)
		key := []byte(config.Session.Key)
		redi_session, err := redistore.NewRediStore(config.Redis.Size, "tcp", address, config.Redis.Pass, key)
		if err != nil {
			logger.Error(err)
			log.Fatalln(err)
		}

		redi_session.Options = &sessions.Options{
			Domain:   config.Session.Host,
			Path:     config.Session.Path,
			MaxAge:   86400 * 256, // 256 Days
			HttpOnly: true,
			Secure:   config.Session.Secure,
			SameSite: http.SameSiteNoneMode,
		}

		session = &Session{redi_session, config}
	})

	return session
}

// Session - Структура хранилища сессии
type Session struct {
	rs     *redistore.RediStore
	config *Config
}

// auth - Сессия авторизации
func (s Session) auth(w http.ResponseWriter, r *http.Request) (*sessions.Session, error) {
	return s.rs.Get(r, authSessionID)
}

// AuthStart - Сессия авторизации. UUID сессии, текущий пользователь и так далее.
func (s *Session) AuthStart(w http.ResponseWriter, r *http.Request) (*models.SessionResponse, error) {
	session, _ := s.auth(w, r)

	// Если мы не видим сесии - то давайте сделаем новую
	if sessionID, ok := session.Values["session_id"].(string); !ok || len(sessionID) == 0 {
		strUUID := uuid.New()
		session.Values["session_id"] = strUUID.String()
		session.Values["struct"] = models.SessionResponse{
			Auth: false,
		}

		if err := session.Save(r, w); err != nil {
			return nil, err
		}
	}

	if sessionResponse, ok := session.Values["struct"].(models.SessionResponse); ok {
		return &sessionResponse, nil
	}

	return nil, nil
}

// AuthCreate - Создание сессии авторизацияя
func (s Session) AuthCreate(w http.ResponseWriter, r *http.Request, subsite *models.User) (*models.SessionResponse, error) {
	session, err := s.auth(w, r)
	if err != nil {
		return nil, err
	}

	sessionResponse := &models.SessionResponse{}
	sessionResponse.Auth = true
	sessionResponse.BindFromSubsite(subsite)

	session.Values["struct"] = sessionResponse

	if err := session.Save(r, w); err != nil {
		return nil, err
	}

	return sessionResponse, nil
}

// AuthGet - Возвращает сессию авторизации
func (s Session) AuthGet(w http.ResponseWriter, r *http.Request) (*models.SessionResponse, error) {
	session, err := s.auth(w, r)
	if err != nil {
		return nil, err
	}

	if sessionResponse, ok := session.Values["struct"].(models.SessionResponse); ok {
		return &sessionResponse, nil
	}

	return nil, nil
}

// Logout - Выход из сессии
func (s *Session) Logout(w http.ResponseWriter, r *http.Request) error {
	session, err := s.auth(w, r)
	if err != nil {
		return err
	}

	postSessionID := formatter.EscapeString(r.FormValue("session_id"))
	currentSessionID := session.Values["session_id"].(string)
	if postSessionID != currentSessionID {
		return errors.New("Ошибка, блядь")
	}

	session.Values["session_id"] = ""
	session.Values["struct"] = &models.SessionResponse{
		Auth: false,
	}

	if err := session.Save(r, w); err != nil {
		return err
	}
	return nil
}
