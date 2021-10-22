type Welcome struct {
	AvatarURL               string              `json:"avatar_url"`
	CommentEditor           CommentEditor       `json:"commentEditor"`
	Cover                   Cover               `json:"cover"`
	Created                 int64               `json:"created"`
	Description             string              `json:"description"`
	Highlight               string              `json:"highlight"`
	ID                      int64               `json:"id"`
	IsAvailableForMessenger bool                `json:"isAvailableForMessenger"`
	IsEnableWriting         bool                `json:"is_enable_writing"`
	IsOnline                bool                `json:"is_online"`
	IsSubscribed            bool                `json:"is_subscribed"`
	IsSubscribedToNewPosts  bool                `json:"is_subscribed_to_new_posts"`
	IsUnsubscribable        bool                `json:"is_unsubscribable"`
	IsVerified              bool                `json:"is_verified"`
	Name                    string              `json:"name"`
	OnlineStatusText        string              `json:"online_status_text"`
	SubscribersAvatars      []SubscribersAvatar `json:"subscribers_avatars"`
	SubscribersCount        int64               `json:"subscribers_count"`
	Type                    int64               `json:"type"`
	URL                     string              `json:"url"`
}

type CommentEditor struct {
	Enabled bool `json:"enabled"`
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
