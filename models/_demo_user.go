type Subsite struct {
	ActiveUntil             int64               `json:"active_until"`
	AdvancedAccess          AdvancedAccess      `json:"advanced_access"`
	Avatar                  Avatar              `json:"avatar"`
	AvatarURL               string              `json:"avatar_url"`
	BannedInfo              []interface{}       `json:"banned_info"`
	CanChangeAvatar         bool                `json:"can_change_avatar"`
	CanChangeCover          bool                `json:"can_change_cover"`
	Counters                Counters            `json:"counters"`
	Cover                   Avatar              `json:"cover"`
	Created                 int64               `json:"created"`
	CreatedRFC              string              `json:"createdRFC"`
	Description             string              `json:"description"`
	ID                      int64               `json:"id"`
	IsAvailableForMessenger bool                `json:"isAvailableForMessenger"`
	IsBanned                bool                `json:"is_banned"`
	IsMuted                 bool                `json:"is_muted"`
	IsPlus                  bool                `json:"is_plus"`
	IsSubscribed            bool                `json:"is_subscribed"`
	IsSubscribedToNewPosts  bool                `json:"is_subscribed_to_new_posts"`
	IsSubsitesTuned         bool                `json:"is_subsites_tuned"`
	IsVerified              bool                `json:"is_verified"`
	Karma                   int64               `json:"karma"`
	MHash                   string              `json:"m_hash"`
	MHashExpirationTime     int64               `json:"m_hash_expiration_time"`
	Name                    string              `json:"name"`
	PushTopic               string              `json:"push_topic"`
	SocialAccounts          []SocialAccount     `json:"social_accounts"`
	SubscribersAvatars      []SubscribersAvatar `json:"subscribers_avatars"`
	SubscribersCount        int64               `json:"subscribers_count"`
	Type                    int64               `json:"type"`
	URL                     string              `json:"url"`
	UserHash                string              `json:"user_hash"`
	UserHashes              []UserHash          `json:"user_hashes"`
}

type AdvancedAccess struct {
	IsNeedsAdvancedAccess bool   `json:"is_needs_advanced_access"`
	Hash                  string `json:"hash"`
}

type Avatar struct {
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

type Counters struct {
	Entries   int64 `json:"entries"`
	Comments  int64 `json:"comments"`
	Favorites int64 `json:"favorites"`
}

type SocialAccount struct {
	ID       int64  `json:"id"`
	Type     int64  `json:"type"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type SubscribersAvatar struct {
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
}

type UserHash struct {
	Integer *int64
	String  *string
}
