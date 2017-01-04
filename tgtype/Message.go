package tgtype

// Message - This object represents a message.
type Message struct {
	MessageID             int32           `json:"message_id"`
	From                  *User           `json:"from,omitempty"`
	Date                  int32           `json:"date"`
	Chat                  *Chat           `json:"chat"`
	ForwardFrom           *User           `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat           `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  int32           `json:"forward_from_message_id,omitempty"`
	ForwardDate           int32           `json:"forward_date,omitempty"`
	ReplyToMessage        *Message        `json:"reply_to_message,omitempty"`
	EditDate              int32           `json:"edit_date,omitempty"`
	Text                  string          `json:"text,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	Audio                 *Audio          `json:"audio,omitempty"`
	Document              *Document       `json:"document,omitempty"`
	Game                  *Game           `json:"game,omitempty"`
	Photo                 []PhotoSize     `json:"photo,omitempty"`
	Sticker               *Sticker        `json:"sticker,omitempty"`
	Video                 *Video          `json:"video,omitempty"`
	Voice                 *Voice          `json:"voice,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	Contact               *Contact        `json:"contact,omitempty"`
	Location              *Location       `json:"location,omitempty"`
	Venue                 *Venue          `json:"venue,omitempty"`
	NewChatMember         *User           `json:"new_chat_member,omitempty"`
	LeftChatMember        *User           `json:"left_chat_member,omitempty"`
	NewChatTitle          string          `json:"new_chat_title,omitempty"`
	NewChatPhoto          []PhotoSize     `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool            `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool            `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool            `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool            `json:"channel_chat_created,omitempty"`
	MigrateToChatID       int32           `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     int32           `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message        `json:"pinned_message,omitempty"`
}
