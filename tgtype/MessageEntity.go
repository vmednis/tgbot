package tgtype

// MessageEntity - This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	TypeString string `json:"type"`
	Offset     int64  `json:"offset"`
	Length     int64  `json:"length"`
	URL        string `json:"url,omitempty"`
	User       *User  `json:"user,omitempty"`
}
