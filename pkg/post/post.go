package post

import "time"

// Post created by a user.
type Post struct {
	ID        uint      `json:"id,omitempty"`
	Title 	  string	`json:"title,omitempty"`
	Category  string	`json:"category,omitempty"`
	Body      string    `json:"body,omitempty"`
	UserID    uint      `json:"user_id,omitempty"`
	SubjectId uint 		`json:"subject_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
