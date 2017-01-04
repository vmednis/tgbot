package tgtype

// Animation - You can provide an animation for your game so that it looks stylish in chats (check out Lumberjack for an example). This object represents an animation file to be displayed in the message containing a game.
type Animation struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int64      `json:"file_size,omitempty"`
}
