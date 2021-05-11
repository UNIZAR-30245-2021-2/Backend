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
func (rr *ReplyRepository) GetByPost(ctx context.Context, postID uint) ([]reply.Reply, error) {
	q := `
	SELECT id, user_id, body, created_at, updated_at
		FROM replies
		WHERE post_id = $1
		ORDER BY created_at;
	`

	rows, err := rr.Data.DB.QueryContext(ctx, q, postID)
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
func (rr *ReplyRepository) Create(ctx context.Context, reply *reply.Reply) error {
	q := `
	INSERT INTO replies (user_id, post_id, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	stmt, err := rr.Data.DB.PrepareContext(ctx, q)
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
func (rr *ReplyRepository) Update(ctx context.Context, id uint, reply reply.Reply) error {
	q := `
	UPDATE replies set body=$1, updated_at=$2
		WHERE id=$3;
	`

	stmt, err := rr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, reply.Body, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a reply by id.
func (rr *ReplyRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM replies WHERE id=$1;`

	stmt, err := rr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
