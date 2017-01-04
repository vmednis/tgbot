package tgtype

// UserProfilePhotos - This object represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int64      `json:"total_count"`
	Photos     *PhotoSize `json:"photos"`
}
