package services

import (
	"context"

	"github.com/gadhittana01/todolist/pkg/domain"
)

type (
	ActivityResource interface {
		GetActivities(ctx context.Context) ([]domain.Activity, error)
		GetActivityByID(ctx context.Context, req domain.GetActivityByID) (domain.Activity, error)
		CreateActivity(ctx context.Context, req domain.CreateActivity) (domain.Activity, error)
		UpdateActivity(ctx context.Context, req domain.UpdateActivity) (domain.Activity, error)
		DeleteActivity(ctx context.Context, req domain.DeleteActivity) (domain.Activity, error)
	}

	TodoResource interface {
		GetTodos(ctx context.Context, activity_group_id int) ([]domain.Todo, error)
		GetTodoByID(ctx context.Context, req domain.GetTodoByID) (domain.Todo, error)
		CreateTodo(ctx context.Context, req domain.CreateTodo) (domain.Todo, error)
		UpdateTodo(ctx context.Context, req domain.UpdateTodo) (domain.Todo, error)
		DeleteTodo(ctx context.Context, req domain.DeleteTodo) (domain.Todo, error)
	}
)
