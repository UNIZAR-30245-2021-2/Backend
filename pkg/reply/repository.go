package reply

import "context"

// Repository handle the CRUD operations with Replies.
type Repository interface {
	GetByPost(ctx context.Context, postID uint) ([]Reply, error)
	Create(ctx context.Context, reply *Reply) error
	Update(ctx context.Context, id uint, reply Reply) error
	Delete(ctx context.Context, id uint) error
}
