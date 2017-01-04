package tgtype

// Video - This object represents a video file.
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int64      `json:"width"`
	Height   int64      `json:"height"`
	Duration int64      `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int64      `json:"file_size,omitempty"`
}
