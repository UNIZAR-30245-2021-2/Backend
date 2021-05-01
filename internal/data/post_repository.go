package data

import (
	"context"
	"time"

	"github.com/orlmonteverde/go-postgres-microblog/pkg/post"
)

// PostRepository manages the operations with the database that
// correspond to the post model.
type PostRepository struct {
	Data *Data
}

// GetAll returns all posts.
func (pr *PostRepository) GetAll(ctx context.Context) ([]post.Post, error) {
	q := `
	SELECT id, title, category, body, user_id, subject_id, created_at, updated_at
		FROM posts;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		rows.Scan(&p.ID, &p.Title, &p.Category, &p.Body, &p.UserID, &p.SubjectId,
			&p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

// GetOne returns one post by id.
func (pr *PostRepository) GetOne(ctx context.Context, id uint) (post.Post, error) {
	q := `
	SELECT id, title, category, body, user_id, subject_id, created_at, updated_at
		FROM posts WHERE id = $1;
	`

	row := pr.Data.DB.QueryRowContext(ctx, q, id)

	var p post.Post
	err := row.Scan(&p.ID, &p.Title, &p.Category, &p.Body, &p.UserID, &p.SubjectId,
		&p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return post.Post{}, err
	}

	return p, nil
}

// GetBySubject returns all subject posts.
func (pr *PostRepository) GetBySubject(ctx context.Context, subjectID uint, order string) ([]post.Post, error) {
	q_created := `
	SELECT id, user_id, title, category, created_at, updated_at
		FROM posts
		WHERE subject_id = $1
		ORDER BY created_at;
	`
	q_updated := `
	SELECT id, user_id, title, category, created_at, updated_at
		FROM posts
		WHERE subject_id = $1
		ORDER BY updated_at;
	`
	var q string
	if order == "created" {
		q = q_created
	} else { // order == "updated"
		q = q_updated
	}

	rows, err := pr.Data.DB.QueryContext(ctx, q, subjectID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Category,
			&p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

// GetByUser returns all user posts.
func (pr *PostRepository) GetByUser(ctx context.Context, userID uint) ([]post.Post, error) {
	q := `
	SELECT id, title, category, body, user_id, subject_id, created_at, updated_at
		FROM posts
		WHERE user_id = $1;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		rows.Scan(&p.ID, &p.Title, &p.Category, &p.Body, &p.UserID, &p.SubjectId,
			&p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) GetByCategory(ctx context.Context, subjectID uint, category string) ([]post.Post, error) {
	q := `
	SELECT id, user_id, title, category, created_at, updated_at
	FROM posts
	WHERE subject_id = $1 AND category LIKE $2;
	`

	rows, err := pr.Data.DB.QueryContext(ctx, q, subjectID, category)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Category, &p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

func (pr *PostRepository) GetByTitle(ctx context.Context, subjectID uint, title string) ([]post.Post, error) {
	q := `
	SELECT id, user_id, title, category, created_at, updated_at
	FROM posts
	WHERE subject_id = $1 AND title LIKE $2;
	`
	title = title + "%"
	rows, err := pr.Data.DB.QueryContext(ctx, q, subjectID, title)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []post.Post
	for rows.Next() {
		var p post.Post
		rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Category, &p.CreatedAt, &p.UpdatedAt)
		posts = append(posts, p)
	}

	return posts, nil
}

// Create adds a new post.
func (pr *PostRepository) Create(ctx context.Context, p *post.Post) error {
	q := `
	INSERT INTO posts (user_id, subject_id, title, category, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.UserID, p.SubjectId, p.Title, p.Category,
		p.Body, time.Now(), time.Now())

	err = row.Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a post by id.
func (pr *PostRepository) Update(ctx context.Context, id uint, p post.Post) error {
	q := `
	UPDATE posts set title=$1, category=$2, body=$3, updated_at=$4
		WHERE id=$5;
	`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Title, p.Category, p.Body, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a post by id.
func (pr *PostRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM posts WHERE id=$1;`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
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
