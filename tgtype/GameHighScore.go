package tgtype

// GameHighScore - This object represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int32 `json:"position"`
	User     *User `json:"user"`
	Score    int32 `json:"score"`
}
