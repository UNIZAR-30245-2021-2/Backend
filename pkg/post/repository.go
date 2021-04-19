package post

import "context"

// Repository handle the CRUD operations with Posts.
type Repository interface {
	GetAll(ctx context.Context) ([]Post, error)
	GetOne(ctx context.Context, id uint) (Post, error)
	GetBySubject(ctx context.Context, subjectID uint) ([]Post, error)
	GetByUser(ctx context.Context, userID uint) ([]Post, error)
	GetByCategory(ctx context.Context, subjectID uint, category string) ([]Post, error)
	GetByTitle(ctx context.Context, subjectID uint, title string) ([]Post, error)
	Create(ctx context.Context, post *Post) error
	Update(ctx context.Context, id uint, post Post) error
	Delete(ctx context.Context, id uint) error
}
