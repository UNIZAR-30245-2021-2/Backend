package data

import (
	"context"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/reply"
	"time"
)

// ReplyRepository manages the operations with the database that
// correspond to the reply model.
type ReplyRepository struct {
	Data *Data
}

// GetByPost returns all post replies.
func (r ReplyRepository) GetByPost(ctx context.Context, postID uint) ([]reply.Reply, error) {
	q := `
	SELECT id, user_id, body, created_at, updated_at
		FROM replies
		WHERE post_id = $1
		ORDER BY created_at;
	`

	rows, err := r.Data.DB.QueryContext(ctx, q, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var replies []reply.Reply
	for rows.Next() {
		var r reply.Reply
		rows.Scan(&r.ID, &r.UserID, &r.Body, &r.CreatedAt, &r.UpdatedAt)
		replies = append(replies, r)
	}

	return replies, nil
}

// Create adds a new reply.
func (r ReplyRepository) Create(ctx context.Context, reply *reply.Reply) error {
	q := `
	INSERT INTO replies (user_id, post_id, body, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	stmt, err := r.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, reply.UserID, reply.PostId, reply.Body,
		time.Now(), time.Now())

	err = row.Scan(&reply.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a reply by id.
func (r ReplyRepository) Update(ctx context.Context, id uint, reply reply.Reply) error {
	panic("implement me")
}

// Delete removes a reply by id.
func (r ReplyRepository) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}
