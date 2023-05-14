package todo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gadhittana01/todolist/pkg/domain"
)

var (
	getTodos = `
		SELECT 
			todo_id,
			activity_group_id,
			title,
			priority,
			is_active,
			created_at,
			updated_at
		FROM todos
	`

	getTodosWithAID = `
		SELECT 
			todo_id,
			activity_group_id,
			title,
			priority,
			is_active,
			created_at,
			updated_at
		FROM todos
		WHERE activity_group_id = ?
	`

	getTodoByID = `
		SELECT 
			todo_id,
			activity_group_id,
			title,
			priority,
			is_active,
			created_at,
			updated_at
		FROM todos
		WHERE todo_id = ?
	`

	insertTodo = `
		INSERT INTO todos(title, activity_group_id, is_active) VALUES (?, ?, ?)
	`

	updateTodo = `
		UPDATE todos
		SET title = ?,
			priority = ?,
			is_active = ?
		WHERE todo_id = ?
	`

	deleteTodo = `
		DELETE FROM todos
		WHERE todo_id = ?
	`
)

type persistent interface {
	// activity
	getTodos(ctx context.Context, activity_group_id int) ([]domain.Todo, error)
	getTodoByID(ctx context.Context, req domain.GetTodoByID) (domain.Todo, error)
	createTodo(ctx context.Context, req domain.CreateTodo) (domain.Todo, error)
	updateTodo(ctx context.Context, req domain.UpdateTodo) (domain.Todo, error)
	deleteTodo(ctx context.Context, req domain.DeleteTodo) (domain.Todo, error)
}

type psql struct {
	db *sql.DB
}

func newPersistent(db *sql.DB) persistent {
	return psql{
		db: db,
	}
}

func (p psql) getTodos(ctx context.Context, activity_group_id int) ([]domain.Todo, error) {
	var res = []domain.Todo{}
	var resQuery *sql.Rows
	var err error

	if activity_group_id == 0 {
		resQuery, err = p.db.Query(getTodos)
		if err != nil {
			return res, err
		}
	} else if activity_group_id > 0 {
		resQuery, err = p.db.Query(getTodosWithAID, activity_group_id)
		if err != nil {
			return res, err
		}
	}

	for resQuery.Next() {
		var id int
		var activity_group_id int
		var title string
		var is_active bool
		var priority sql.NullString
		var createdAt string
		var updatedAt sql.NullString

		err = resQuery.Scan(&id, &activity_group_id, &title, &priority, &is_active, &createdAt, &updatedAt)
		if err != nil {
			return res, err
		}

		res = append(res, domain.Todo{
			ID:              id,
			ActivityGroupID: activity_group_id,
			Title:           title,
			IsActive:        is_active,
			Priority:        priority.String,
			CreateAt:        createdAt,
			UpdatedAt:       updatedAt.String,
		})
	}

	return res, nil
}

func (p psql) getTodoByID(ctx context.Context, req domain.GetTodoByID) (domain.Todo, error) {
	var res = domain.Todo{}
	var id int
	var activity_group_id int
	var title string
	var is_active bool
	var priority sql.NullString
	var createdAt string
	var updatedAt sql.NullString

	err := p.db.QueryRow(getTodoByID, req.ID).Scan(&id, &activity_group_id, &title, &priority, &is_active, &createdAt, &updatedAt)
	if err != nil {
		return res, err
	}

	res = domain.Todo{
		ID:              id,
		ActivityGroupID: activity_group_id,
		Title:           title,
		IsActive:        is_active,
		Priority:        priority.String,
		CreateAt:        createdAt,
		UpdatedAt:       updatedAt.String,
	}
	return res, nil
}

func (p psql) createTodo(ctx context.Context, req domain.CreateTodo) (domain.Todo, error) {
	var res = domain.Todo{}

	resQuery, err := p.db.ExecContext(ctx, insertTodo, req.Title, req.ActivityGroupID, req.IsActive)
	if err != nil {
		return res, err
	}

	lastID, err := resQuery.LastInsertId()
	if err != nil {
		return res, err
	}

	return p.getTodoByID(ctx, domain.GetTodoByID{
		ID: int(lastID),
	})
}

func (p psql) updateTodo(ctx context.Context, req domain.UpdateTodo) (domain.Todo, error) {
	var res = domain.Todo{}

	_, err := p.db.ExecContext(ctx, updateTodo, req.Title, req.Priority, req.IsActive, req.ID)
	if err != nil {
		return res, err
	}

	return p.getTodoByID(ctx, domain.GetTodoByID{
		ID: req.ID,
	})
}

func (p psql) deleteTodo(ctx context.Context, req domain.DeleteTodo) (domain.Todo, error) {
	var res = domain.Todo{}

	resQuery, err := p.db.ExecContext(ctx, deleteTodo, req.ID)
	if err != nil {
		return res, err
	}

	if rows, _ := resQuery.RowsAffected(); rows == 0 {
		return res, errors.New("sql: no rows in result set")
	}

	return res, nil
}
