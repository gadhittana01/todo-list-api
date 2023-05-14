package todo

import (
	"context"
	"database/sql"

	"github.com/gadhittana01/todolist/config"
	"github.com/gadhittana01/todolist/pkg/domain"
)

type IResource interface {
	GetTodos(ctx context.Context, activity_group_id int) ([]domain.Todo, error)
	GetTodoByID(ctx context.Context, req domain.GetTodoByID) (domain.Todo, error)
	CreateTodo(ctx context.Context, req domain.CreateTodo) (domain.Todo, error)
	UpdateTodo(ctx context.Context, req domain.UpdateTodo) (domain.Todo, error)
	DeleteTodo(ctx context.Context, req domain.DeleteTodo) (domain.Todo, error)
}

type module struct {
	persistent persistent
}

func New(cfg *config.GlobalConfig, db *sql.DB) IResource {
	return &module{
		persistent: newPersistent(db),
	}
}

func (m module) GetTodos(ctx context.Context, activity_group_id int) ([]domain.Todo, error) {
	return m.persistent.getTodos(ctx, activity_group_id)
}

func (m module) GetTodoByID(ctx context.Context, req domain.GetTodoByID) (domain.Todo, error) {
	return m.persistent.getTodoByID(ctx, req)
}

func (m module) CreateTodo(ctx context.Context, req domain.CreateTodo) (domain.Todo, error) {
	return m.persistent.createTodo(ctx, req)
}

func (m module) UpdateTodo(ctx context.Context, req domain.UpdateTodo) (domain.Todo, error) {
	return m.persistent.updateTodo(ctx, req)
}

func (m module) DeleteTodo(ctx context.Context, req domain.DeleteTodo) (domain.Todo, error) {
	return m.persistent.deleteTodo(ctx, req)
}
