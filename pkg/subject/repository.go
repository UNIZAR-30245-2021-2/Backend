package subject

import "context"

// Repository handle the CRUD operations with Subjects.
type Repository interface {
	GetAll(ctx context.Context) ([]Subject, error)
	GetOne(ctx context.Context, id uint) (Subject, error)
	GetByYear(ctx context.Context, year int) ([]Subject, error)
	Create(ctx context.Context, subject *Subject) error
	Update(ctx context.Context, id uint, subject Subject) error
	Delete(ctx context.Context, id uint) error
}
