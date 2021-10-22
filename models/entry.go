package models

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/yuriygr/go-mlh/formatter"
	"github.com/yuriygr/go-posledstvie/utils"
)

// Entry
type Entry struct {
	EntryID uint32 `json:"id"    db:"e.entry_id"`
	Slug    string `json:"slug"`

	Autor   EntryAuthor  `json:"author"  db:"author"`
	Subsite EntrySubsite `json:"subsite" db:"subsite"`

	Content struct {
		Title  string `json:"title"   db:"title"`
		Intro  string `json:"intro"   db:"intro"`
		Verion uint8  `json:"version" db:"version"`
	} `json:"content"  db:"ec"`

	Counters EntryCounters `json:"counters"  db:"counters"`

	Vote *string `json:"vote" db:"vote"`

	States struct {
		IsEnabledComments      bool `json:"is_enabled_comments"        db:"e.is_enabled_comments"`
		IsEnabledLikes         bool `json:"is_enabled_likes"           db:"e.is_enabled_likes"`
		IsFavorited            bool `json:"is_favorited"               db:"is_favorited"`
		IsReposted             bool `json:"is_reposted"                db:"is_reposted"`
		IsPinned               bool `json:"is_pinned"                  db:"e.is_pinned"`
		IsShowThanks           bool `json:"is_show_thanks"             db:"e.is_show_thanks"`
		IsStillUpdating        bool `json:"is_still_updating"          db:"e.is_still_updating"`
		IsFilledByEditors      bool `json:"is_filled_by_editors"       db:"e.is_filled_by_editors"`
		IsEditorial            bool `json:"is_editorial"               db:"e.is_editorial"`
		IsDraft                bool `json:"is_draft"                   db:"e.is_draft"`
		IsSubscribedToNewPosts bool `json:"is_subscribed_to_new_posts" db:"is_subscribed_to_new_posts"`
	} `json:"states"  db:""`

	Created string `json:"created" db:"e.created"`
}

// Render - Make HTTP status code equal to status code in struct
func (sr *Entry) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// NewEntryResponse -
func NewEntryResponse(items []*Entry) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range items {
		list = append(list, item)
	}
	return list
}

// EntryAuthor -
type EntryAuthor struct {
	ID     int64  `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`

	IsMe                   JsonBool `json:"is_me"`
	IsPlus                 JsonBool `json:"is_plus"`
	IsUser                 JsonBool `json:"is_user"`
	IsOnline               JsonBool `json:"is_online"`
	IsVerified             JsonBool `json:"is_verified"`
	IsSubscribed           JsonBool `json:"is_subscribed"`
	IsSubscribedToNewPosts JsonBool `json:"is_subscribed_to_new_posts"`
}

// EntrySubsite -
type EntrySubsite struct {
	ID     int64  `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`

	IsVerified             JsonBool `json:"is_verified"`
	IsEnableWriting        JsonBool `json:"is_enable_writing"`
	IsSubscribed           JsonBool `json:"is_subscribed"`
	IsSubscribedToNewPosts JsonBool `json:"is_subscribed_to_new_posts"`
}

// EntryCounters - Названия отдельной структурой для имплементации методов сканирования
type EntryCounters struct {
	Votes    int16  `json:"votes"`
	Comments uint16 `json:"comments"`
	Reposts  uint16 `json:"reposts"`
	Hits     uint16 `json:"hits"`
}

// EntriesRequest -
type EntriesRequest struct {
	CurrentUserID uint32 `db:"r.current_user_id"`

	Sort   string
	Order  string `db:"r.order"`
	Offset uint32

	SubsiteID   uint32 `db:"r.subsite_id"`
	PinnedOnTop bool
}

// Bind - Bind HTTP request data and validate it
func (pr *EntriesRequest) Bind(r *http.Request) error {

	//avalibleSorts := []string{"/all", "/all/new", "/new", "/new/new"}
	//if sort := r.URL.Query().Get("sort"); sort != "" {
	//	sort = formatter.EscapeString(sort)
	//	if helpers.StringInSlice(sort, avalibleSorts) {
	//		pr.Sort = sort
	//	}
	//}

	//if order := r.URL.Query().Get("order"); order != "" {
	//	pr.Order = formatter.EscapeString(order)
	//}

	//if offset := r.URL.Query().Get("offset"); offset != "" {
	//	pr.Offset = utils.Uint32(offset)
	//}

	if chi.URLParam(r, "param") != "" {
		pr.SubsiteID = utils.Uint32(formatter.EscapeString(chi.URLParam(r, "param")))
	}

	return nil
}

// EntryRequest -
type EntryRequest struct {
	CurrentUserID uint32 `db:"r.current_user_id"`
	EntryID       uint32 `db:"r.entry_id"`
}

// Bind - Bind HTTP request data and validate it
func (pr *EntryRequest) Bind(r *http.Request) error {
	pr.EntryID = utils.Uint32(formatter.EscapeString(chi.URLParam(r, "id")))

	return nil
}

// EntryLikeRequest -
type EntryLikeRequest struct {
	CurrentUserID uint32 `db:"r.current_user_id"`
	EntryID       uint32 `db:"r.entry_id"`
	Action        string `db:"r.action"`
}

// Bind - Bind HTTP request data and validate it
func (pr *EntryLikeRequest) Bind(r *http.Request) error {
	return nil
}
