package tgtype

// Sticker - This object represents a sticker.
type Sticker struct {
	FileID   string     `json:"file_id"`
	Width    int32      `json:"width"`
	Height   int32      `json:"height"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	Emoji    string     `json:"emoji,omitempty"`
	FileSize int32      `json:"file_size,omitempty"`
}
