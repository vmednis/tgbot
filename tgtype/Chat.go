package tgtype

// Chat - This object represents a chat.
type Chat struct {
	ID                          int64  `json:"id"`
	TypeString                  string `json:"type"`
	Title                       string `json:"title, omitempty"`
	Username                    string `json:"username, omitempty"`
	FirstName                   string `json:"first_name, omitempty"`
	LastName                    string `json:"last_name, omitempty"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators, omitempty"`
}
