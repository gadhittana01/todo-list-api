package activity

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gadhittana01/todolist/pkg/domain"
)

var (
	getActivities = `
		SELECT 
			activity_id,
			title,
			email,
			created_at,
			updated_at
		FROM activities
	`

	getActivitiesByID = `
		SELECT 
			activity_id,
			title,
			email,
			created_at,
			updated_at
		FROM activities
		WHERE activity_id = ?
	`

	insertActivity = `
		INSERT INTO activities(title, email) VALUES (?, ?)
	`

	updateActivity = `
		UPDATE activities 
		SET title = ?
		WHERE activity_id = ?
	`

	deleteActivity = `
		DELETE FROM activities 
		WHERE activity_id = ?
	`
)

type persistent interface {
	// activity
	getActivities(ctx context.Context) ([]domain.Activity, error)
	getActivityByID(ctx context.Context, req domain.GetActivityByID) (domain.Activity, error)
	createActivity(ctx context.Context, req domain.CreateActivity) (domain.Activity, error)
	updateActivity(ctx context.Context, req domain.UpdateActivity) (domain.Activity, error)
	deleteActivity(ctx context.Context, req domain.DeleteActivity) (domain.Activity, error)
}

type psql struct {
	db *sql.DB
}

func newPersistent(db *sql.DB) persistent {
	return psql{
		db: db,
	}
}

func (p psql) getActivities(ctx context.Context) ([]domain.Activity, error) {
	var res = []domain.Activity{}

	resQuery, err := p.db.Query(getActivities)
	if err != nil {
		return res, err
	}

	for resQuery.Next() {
		var id int
		var title string
		var email string
		var createdAt string
		var updatedAt sql.NullString

		err = resQuery.Scan(&id, &title, &email, &createdAt, &updatedAt)
		if err != nil {
			return res, err
		}

		res = append(res, domain.Activity{
			ID:        id,
			Title:     title,
			Email:     email,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt.String,
		})
	}

	return res, nil
}

func (p psql) getActivityByID(ctx context.Context, req domain.GetActivityByID) (domain.Activity, error) {
	var res = domain.Activity{}
	var id int
	var title string
	var email string
	var createdAt string
	var updatedAt sql.NullString

	err := p.db.QueryRow(getActivitiesByID, req.ID).Scan(&id, &title, &email, &createdAt, &updatedAt)
	if err != nil {
		return res, err
	}

	res = domain.Activity{
		ID:        id,
		Title:     title,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt.String,
	}
	return res, nil
}

func (p psql) createActivity(ctx context.Context, req domain.CreateActivity) (domain.Activity, error) {
	var res = domain.Activity{}

	resQuery, err := p.db.ExecContext(ctx, insertActivity, req.Title, req.Email)
	if err != nil {
		return res, err
	}

	lastID, err := resQuery.LastInsertId()
	if err != nil {
		return res, err
	}

	return p.getActivityByID(ctx, domain.GetActivityByID{
		ID: int(lastID),
	})
}

func (p psql) updateActivity(ctx context.Context, req domain.UpdateActivity) (domain.Activity, error) {
	var res = domain.Activity{}

	_, err := p.db.ExecContext(ctx, updateActivity, req.Title, req.ID)
	if err != nil {
		return res, err
	}

	return p.getActivityByID(ctx, domain.GetActivityByID{
		ID: req.ID,
	})
}

func (p psql) deleteActivity(ctx context.Context, req domain.DeleteActivity) (domain.Activity, error) {
	var res = domain.Activity{}

	resQuery, err := p.db.ExecContext(ctx, deleteActivity, req.ID)
	if err != nil {
		return res, err
	}

	if rows, _ := resQuery.RowsAffected(); rows == 0 {
		return res, errors.New("sql: no rows in result set")
	}

	return res, nil
}
