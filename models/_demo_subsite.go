type Subsite struct {
	AvatarURL               string              `json:"avatar_url"`
	CommentEditor           CommentEditor       `json:"commentEditor"`
	CommentsCount           int64               `json:"comments_count"`
	Contacts                Contacts            `json:"contacts"`
	Cover                   Cover               `json:"cover"`
	Created                 int64               `json:"created"`
	CreatedRFC              string              `json:"createdRFC"`
	Description             string              `json:"description"`
	EntriesCount            int64               `json:"entries_count"`
	EventsCount             int64               `json:"events_count"`
	ID                      int64               `json:"id"`
	IsAvailableForMessenger bool                `json:"isAvailableForMessenger"`
	IsEnableWriting         bool                `json:"is_enable_writing"`
	IsMuted                 bool                `json:"is_muted"`
	IsSubscribed            bool                `json:"is_subscribed"`
	IsSubscribedToNewPosts  bool                `json:"is_subscribed_to_new_posts"`
	IsUnsubscribable        bool                `json:"is_unsubscribable"`
	IsVerified              bool                `json:"is_verified"`
	Karma                   int64               `json:"karma"`
	Name                    string              `json:"name"`
	Rules                   string              `json:"rules"`
	SubscribersAvatars      []SubscribersAvatar `json:"subscribers_avatars"`
	SubscribersCount        int64               `json:"subscribers_count"`
	Type                    int64               `json:"type"`
	URL                     string              `json:"url"`
	VacanciesCount          int64               `json:"vacancies_count"`
}

type CommentEditor struct {
	Enabled bool `json:"enabled"`
}

type Contacts struct {
	Socials  []interface{} `json:"socials"`
	Site     []interface{} `json:"site"`
	Email    string        `json:"email"`
	Contacts string        `json:"contacts"`
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

type SubscribersAvatar struct {
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
}
