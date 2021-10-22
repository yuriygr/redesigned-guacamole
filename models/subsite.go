package models

import (
	"net/http"

	"github.com/go-chi/render"
)

// MiniSubsite - Подсайт в миниатюре. Только для списков
type MiniSubsite struct {
	ID   int64  `json:"id"          db:"s.subsite_id"`
	Name string `json:"label"       db:"s.name"`
	Slug string `json:"slug"        db:"s.slug"`

	// Думаю надо отдельную таблицу для такого с вызовом через сохраненную функцию
	Avatar string `json:"image"   db:"s.avatar"`
	Cover  string `json:"cover"   db:"s.cover"`

	IsPlus       bool `json:"is_plus"        db:"s.is_plus"`
	IsPlusHidden bool `json:"is_plus_hidden" db:"s.is_plus_hidden"`

	IsMe bool `json:"is_me" db:"is_me"`

	IsSubscribed bool `json:"is_subscribed"  db:"is_subscribed"`

	IsUser     bool  `json:"is_user"        db:"s.is_user"`
	IsVerified bool  `json:"is_verified"    db:"s.is_verified"`
	Status     uint8 `json:"-"              db:"s.status"`
}

// Render - Render, wtf
func (s *MiniSubsite) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewMiniSubsResponse - Вывод объекта списком рендера
func NewMiniSubsResponse(subs []*MiniSubsite) []render.Renderer {
	list := []render.Renderer{}
	for _, sub := range subs {
		list = append(list, sub)
	}
	return list
}

// BaseSubsite - Базовая структура подсайта
type BaseSubsite struct {
	ID          uint32 `json:"id"          db:"s.subsite_id"`
	Name        string `json:"name"        db:"s.name"`
	Slug        string `json:"slug"        db:"s.slug"`
	Description string `json:"description" db:"s.description"`
	Karma       int64  `json:"karma"       db:"s.karma"`

	// Думаю надо отдельную таблицу для такого с вызовом через сохраненную функцию
	Avatar string `json:"avatar" db:"s.avatar"`
	Cover  string `json:"cover"  db:"s.cover"`
	//Avatar Avatar `json:"avatar"`
	//Cover Cover `json:"cover"`

	// Штуки, которые расчитываются исходя из совпадения ID подсайта и текущего ID пользователя и доступов модерации
	CanChangeAvatar bool `json:"can_change_avatar"`
	CanChangeCover  bool `json:"can_change_cover"`

	IsAvailableForMessenger bool `json:"is_available_for_messenger" db:"is_available_for_messenger"`
	IsMuted                 bool `json:"is_muted"                   db:"s.is_muted"`

	IsMe bool `json:"is_me" db:"is_me"`

	IsSubscribed           bool `json:"is_subscribed"              db:"is_subscribed"`
	IsSubscribedToNewPosts bool `json:"is_subscribed_to_new_posts" db:"is_subscribed_to_new_posts"`

	IsUser     bool   `json:"is_user"     db:"s.is_user"`
	IsVerified bool   `json:"is_verified" db:"s.is_verified"`
	Created    string `json:"created"     db:"s.created"`
	Status     uint8  `json:"-"           db:"s.status"`

	Subscribers   SubsiteSubscribers   `json:"subscribers"   db:"subscribers"`
	Subscriptions SubsiteSubscriptions `json:"subscriptions" db:"subscriptions"`

	Contacts struct {
		Socials  []SocialAccount `json:"socials"`
		Site     string          `json:"site"`
		Email    string          `json:"email"`
		Contacts string          `json:"contacts"`
	} `json:"contacts"`

	Rules SubsiteRules `json:"rules" db:"rules"`
}

// User - Структура подсайта-пользователя
type User struct {
	BaseSubsite

	IsBanned     bool `json:"is_banned"      db:"s.is_banned"`
	IsPlus       bool `json:"is_plus"        db:"s.is_plus"`
	IsPlusHidden bool `json:"is_plus_hidden" db:"s.is_plus_hidden"`

	Counters struct {
		Entries   int64 `json:"entries"`
		Comments  int64 `json:"comments"`
		Favorites int64 `json:"favorites"`
	} `json:"counters"`

	Stats        []struct{}
	Schievements []struct{}
}

// Render - Render, wtf
func (s *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Subsite - Структура подсайта
type Subsite struct {
	BaseSubsite

	IsEnableWriting  bool `json:"is_enable_writing" db:"s.is_enable_writing"`
	IsUnsubscribable bool `json:"is_unsubscribable" db:"s.is_unsubscribable"`

	Counters struct {
		Entries   int64 `json:"entries"`
		Comments  int64 `json:"comments"`
		Events    int64 `json:"events"`
		Products  int64 `json:"products"`
		Vacancies int64 `json:"vacancies"`
	} `json:"counters"`
}

// Render - Render, wtf
func (s *Subsite) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Mess

type Avatar struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Cover struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}

type Data struct {
	UUID            string        `json:"uuid"`
	Width           int64         `json:"width"`
	Height          int64         `json:"height"`
	Size            int64         `json:"size"`
	Type            string        `json:"type"`
	Color           string        `json:"color"`
	Hash            string        `json:"hash"`
	ExternalService []interface{} `json:"external_service"`
}

type SocialAccount struct {
	ID       int64  `json:"id"`
	Type     int64  `json:"type"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

// NewSubsiteResponse - Вывод объекта списком рендера
func NewSubsiteResponse(subs []*Subsite) []render.Renderer {
	list := []render.Renderer{}
	for _, sub := range subs {
		list = append(list, sub)
	}
	return list
}

// SubsiteSubscribers -
type SubsiteSubscribers struct {
	Items []SubscribersAvatar `json:"items"`
	Count uint8               `json:"count"`
}

// SubsiteSubscriptions -
type SubsiteSubscriptions struct {
	Items []SubscribersAvatar `json:"items"`
	Count uint8               `json:"count"`
}

// SubsiteRules -
type SubsiteRules struct {
	Items []struct {
		Label string `json:"label"`
	} `json:"items"`
	Count uint8 `json:"count"`
}

// SubscribersAvatar -
type SubscribersAvatar struct {
	ID     uint64   `json:"id"`
	Slug   string   `json:"slug"`
	Image  string   `json:"image"`
	Label  string   `json:"label"`
	IsUser JsonBool `json:"is_user"`
}
