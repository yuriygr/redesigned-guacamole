package models

// Comment -
type Comment struct {
	CommentID uint32 `db:"c.comment_id"`
	ParentID  uint32 `db:"c.paren_id"`
	SubsiteID uint32 `db:"c.subsite_id"`
	EntryID   uint32 `db:"c.entry_id"`

	//Vote

	States struct {
		IsFavorited bool `json:"is_favorited"  db:"is_favorited"`
		IsLiked     bool `json:"is_liked"      db:"is_liked"`
		IsPinned    bool `json:"is_pinned"     db:"c.is_pinned"`
		IsIgnored   bool `json:"is_ignored"    db:"is_ignored"`
		IsRemoved   bool `json:"is_removed"    db:"is_removed"`
		IsEdited    bool `json:"is_edited"     db:"c.is_edited"`
	} `json:"states"  db:""`

	Created string `db:"c.created"`
}
