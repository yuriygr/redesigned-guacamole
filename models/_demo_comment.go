type J struct {
	Message string   `json:"message"`
	Result  []Result `json:"result"`
}

type Result struct {
	ID      int64   `json:"id"`
	Author  Author  `json:"author"`
	ReplyTo int64   `json:"replyTo"`
	Date    int64   `json:"date"`
	Media   []Media `json:"media"`
	Level   int64   `json:"level"`

	Likes Likes `json:"likes"`

	IsPinned    bool `json:"is_pinned"`
	IsIgnored   bool `json:"is_ignored"`
	IsRemoved   bool `json:"is_removed"`
	IsEdited    bool `json:"isEdited"`
	IsFavorited bool `json:"isFavorited"`

	Text        string      `json:"text"`
	TextWoMd    string      `json:"text_wo_md"`
	HTML        string      `json:"html"`
	Attaches    []Attach    `json:"attaches"`
	SourceID    int64       `json:"source_id"`
	Entry       interface{} `json:"entry"`
	LoadMore    LoadMore    `json:"load_more"`
	EtcControls EtcControls `json:"etcControls"`
	Highlight   string      `json:"highlight"`
	Donate      interface{} `json:"donate"`
}

type Attach struct {
	Type Type       `json:"type"`
	Data AttachData `json:"data"`
}

type AttachData struct {
	UUID            *string       `json:"uuid,omitempty"`
	Width           *int64        `json:"width,omitempty"`
	Height          *int64        `json:"height,omitempty"`
	Size            *int64        `json:"size,omitempty"`
	Type            *string       `json:"type,omitempty"`
	Color           *string       `json:"color,omitempty"`
	Hash            *string       `json:"hash,omitempty"`
	ExternalService []interface{} `json:"external_service,omitempty"`
	URL             *string       `json:"url,omitempty"`
	Title           *string       `json:"title,omitempty"`
	Description     *string       `json:"description,omitempty"`
	Image           *ImageClass   `json:"image,omitempty"`
	V               *int64        `json:"v,omitempty"`
}

type ImageClass struct {
	Type Type      `json:"type"`
	Data ImageData `json:"data"`
}

type ImageData struct {
	UUID            string        `json:"uuid"`
	Width           int64         `json:"width"`
	Height          int64         `json:"height"`
	Size            int64         `json:"size"`
	Type            string        `json:"type"`
	Color           string        `json:"color"`
	Hash            string        `json:"hash"`
	ExternalService []interface{} `json:"external_service"`
}

type Author struct {
	ID               int64            `json:"id"`
	Name             string           `json:"name"`
	AvatarURL        string           `json:"avatar_url"`
	IsVerified       bool             `json:"is_verified"`
	Type             int64            `json:"type"`
	IsOnline         bool             `json:"is_online"`
	OnlineStatusText OnlineStatusText `json:"online_status_text"`
}

type EtcControls struct {
	PinComment   bool `json:"pin_comment"`
	Remove       bool `json:"remove"`
	RemoveThread bool `json:"remove_thread"`
}

type Likes struct {
	IsLiked int64 `json:"is_liked"`
	Count   int64 `json:"count"`
	Summ    int64 `json:"summ"`
}

type LoadMore struct {
	Count   int64         `json:"count"`
	IDS     []interface{} `json:"ids"`
	Avatars []interface{} `json:"avatars"`
}

type Media struct {
	Type           int64          `json:"type"`
	ImageURL       string         `json:"imageUrl"`
	AdditionalData AdditionalData `json:"additionalData"`
	Size           Size           `json:"size"`
}

type AdditionalData struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	HasAudio bool   `json:"hasAudio"`
}

type Size struct {
	Width  int64   `json:"width"`
	Height int64   `json:"height"`
	Ratio  float64 `json:"ratio"`
}

type Type string

const (
	Image Type = "image"
	Link  Type = "link"
)

type OnlineStatusText string

const (
	Был11МинутНазад OnlineStatusText = "Был 11 минут назад"
	Был32МинутНазад OnlineStatusText = "Был 32 минут назад"
	Был4МинутНазад  OnlineStatusText = "Был 4 минут назад"
	Был8МинутНазад  OnlineStatusText = "Был 8 минут назад"
	БылНедавно      OnlineStatusText = "Был недавно"
	Онлайн          OnlineStatusText = "Онлайн"
)
