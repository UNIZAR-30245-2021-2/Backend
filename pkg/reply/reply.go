package reply

import "time"

// Reply created by a user.
type Reply struct {
	ID        uint      `json:"id,omitempty"`
	Body      string    `json:"body,omitempty"`
	UserID    uint      `json:"user_id,omitempty"`
	PostId    uint 		`json:"post_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
