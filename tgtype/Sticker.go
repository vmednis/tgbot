package tgtype

// Sticker - This object represents a sticker.
type Sticker struct {
	FileID   string     `json:"file_id"`
	Width    int64      `json:"width"`
	Height   int64      `json:"height"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	Emoji    string     `json:"emoji,omitempty"`
	FileSize int64      `json:"file_size,omitempty"`
}
