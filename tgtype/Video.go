package tgtype

// Video - This object represents a video file.
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int32      `json:"width"`
	Height   int32      `json:"height"`
	Duration int32      `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int32      `json:"file_size,omitempty"`
}
