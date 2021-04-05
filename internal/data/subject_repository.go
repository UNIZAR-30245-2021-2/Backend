package data

import (
	"context"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/subject"
)

// SubjectRepository manages the operations with the database that
// correspond to the subject model.
type SubjectRepository struct {
	Data *Data
}

// GetAll returns all subjects.
func (sr *SubjectRepository) GetAll(ctx context.Context) ([]subject.Subject, error) {
	q := `
	SELECT id, name, year
		FROM subjects;
	`

	rows, err := sr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subjects []subject.Subject
	for rows.Next() {
		var s subject.Subject
		rows.Scan(&s.ID, &s.Name, &s.Year)
		subjects = append(subjects, s)
	}

	return subjects, nil
}

// GetOne returns one subject by id.
func (sr *SubjectRepository) GetOne(ctx context.Context, id uint) (subject.Subject, error) {
	q := `
	SELECT id, name, year
		FROM subjects WHERE id = $1;
	`

	row := sr.Data.DB.QueryRowContext(ctx, q, id)

	var s subject.Subject
	err := row.Scan(&s.ID, &s.Name, &s.Year)
	if err != nil {
		return subject.Subject{}, err
	}

	return s, nil
}

// GetByYear returns all years subjects.
func (sr *SubjectRepository) GetByYear(ctx context.Context, year int) ([]subject.Subject, error) {
	q := `
	SELECT id, name, year
		FROM subjects
		WHERE year = $1;
	`

	rows, err := sr.Data.DB.QueryContext(ctx, q, year)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subjects []subject.Subject
	for rows.Next() {
		var s subject.Subject
		rows.Scan(&s.ID, &s.Name, &s.Year)
		subjects = append(subjects, s)
	}

	return subjects, nil
}

// Create adds a new subject.
func (sr *SubjectRepository) Create(ctx context.Context, s *subject.Subject) error {
	q := `
	INSERT INTO subjects (name, year)
		VALUES ($1, $2)
		RETURNING id;
	`

	stmt, err := sr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, s.Name, s.Year)

	err = row.Scan(&s.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a subject by id.
func (sr *SubjectRepository) Update(ctx context.Context, id uint, s subject.Subject) error {
	q := `
	UPDATE subjects set name=$1, year=$2
		WHERE id=$3;
	`

	stmt, err := sr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, s.Name, s.Year, id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a subject by id.
func (sr *SubjectRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM subjects WHERE id=$1;`

	stmt, err := sr.Data.DB.PrepareContext(ctx, q)
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
