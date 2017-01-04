package tgtype

// PhotoSize - This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int32  `json:"width"`
	Height   int32  `json:"height"`
	FileSize int32  `json:"file_size,omitempty"`
}
