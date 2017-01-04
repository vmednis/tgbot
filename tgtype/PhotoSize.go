package tgtype

// PhotoSize - This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
	FileSize int64  `json:"file_size,omitempty"`
}
