package models

import (
	"net/http"
)

// SessionResponse - Состояние юзера
type SessionResponse struct {
	Auth bool `json:"auth"`

	ID     uint32 `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`

	States struct {
		IsAdult            bool `json:"is_adult"`
		IsEnabledMessenger bool `json:"is_enabled_messenger"`
		IsSubsitesTuned    bool `json:"is_subsites_tuned"`
		IsSuperuser        bool `json:"is_superuser"`
		IsBanned           bool `json:"is_banned"`
		IsUser             bool `json:"is_user"`
		IsKnownEmail       bool `json:"is_known_email"`
		IsPaid             bool `json:"is_paid"`
	} `json:"states"`

	PaidFeatures struct {
		HideBanners    string `json:"hide_banners"`
		HideLiveStream string `json:"hide_live_stream"`
		StraightAngles string `json:"straight_angles"`
	} `json:"paid_features"`

	Notifications struct {
		Count uint32 `json:"count"`
	} `json:"notifications"`

	Created string `json:"created"`

	//CanMarkNsfw    bool        `json:"can_mark_nsfw"`
	//Level          int64       `json:"level"`
	//PaidTillDate   interface{} `json:"paid_till_date"`
	//ShowPossession bool        `json:"show_possession"`
	//StatGroup      string      `json:"stat_group"`
}

// Render - Make HTTP status code equal to status code in struct
func (sr *SessionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind - Bind structure with session
func (sr *SessionResponse) BindFromSubsite(subsite *User) {
	sr.ID = subsite.ID
	sr.Slug = subsite.Slug
	sr.Name = subsite.Name
	sr.Avatar = subsite.Avatar

	sr.States.IsAdult = false
	sr.States.IsEnabledMessenger = false
	sr.States.IsSubsitesTuned = false
	sr.States.IsSuperuser = false
	sr.States.IsBanned = false
	sr.States.IsUser = true
	sr.States.IsKnownEmail = true
	sr.States.IsPaid = false

}
