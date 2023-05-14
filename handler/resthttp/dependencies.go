package resthttp

import (
	"context"

	"github.com/gadhittana01/todolist/services"
)

type (
	ActivityService interface {
		GetActivities(ctx context.Context) ([]services.Activity, error)
		GetActivityByID(ctx context.Context, req services.GetActivityByID) (services.Activity, error)
		CreateActivity(ctx context.Context, req services.CreateActivity) (services.Activity, error)
		UpdateActivity(ctx context.Context, req services.UpdateActivity) (services.Activity, error)
		DeleteActivity(ctx context.Context, req services.DeleteActivity) (services.Activity, error)
	}

	TodoService interface {
		GetTodos(ctx context.Context, activity_group_id int) ([]services.Todo, error)
		GetTodoByID(ctx context.Context, req services.GetTodoByID) (services.Todo, error)
		CreateTodo(ctx context.Context, req services.CreateTodo) (services.Todo, error)
		UpdateTodo(ctx context.Context, req services.UpdateTodo) (services.Todo, error)
		DeleteTodo(ctx context.Context, req services.DeleteTodo) (services.Todo, error)
	}
)
