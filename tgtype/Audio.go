package tgtype

// Audio - This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int64  `json:"duration"`
	Performer string `json:"performer,omitempty"`
	Title     string `json:"title,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
	FileSize  int64  `json:"file_size,omitempty"`
}
