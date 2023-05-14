package services

import (
	"context"

	"github.com/gadhittana01/todolist/pkg/domain"
)

type TodoService interface {
	GetTodos(ctx context.Context, activity_group_id int) ([]Todo, error)
	GetTodoByID(ctx context.Context, req GetTodoByID) (Todo, error)
	CreateTodo(ctx context.Context, req CreateTodo) (Todo, error)
	UpdateTodo(ctx context.Context, req UpdateTodo) (Todo, error)
	DeleteTodo(ctx context.Context, req DeleteTodo) (Todo, error)
}

type todoService struct {
	tr TodoResource
}

func NewTodoService(dep TodoDependencies) (TodoService, error) {
	return &todoService{
		tr: dep.TR,
	}, nil
}

func (p todoService) GetTodos(ctx context.Context, activity_group_id int) ([]Todo, error) {
	var result = []Todo{}

	res, err := p.tr.GetTodos(ctx, activity_group_id)
	if err != nil {
		return result, err
	}

	for _, item := range res {
		result = append(result, Todo{
			ID:              item.ID,
			ActivityGroupID: item.ActivityGroupID,
			Title:           item.Title,
			IsActive:        item.IsActive,
			Priority:        item.Priority,
			CreateAt:        item.CreateAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}

	return result, nil
}

func (p todoService) GetTodoByID(ctx context.Context, req GetTodoByID) (Todo, error) {
	var result = Todo{}

	res, err := p.tr.GetTodoByID(ctx, domain.GetTodoByID{
		ID: req.ID,
	})
	if err != nil {
		return result, err
	}

	result = Todo(res)

	return result, nil
}

func (p todoService) CreateTodo(ctx context.Context, req CreateTodo) (Todo, error) {
	var result = Todo{}

	res, err := p.tr.CreateTodo(ctx, domain.CreateTodo{
		Title:           req.Title,
		ActivityGroupID: req.ActivityGroupID,
		IsActive:        req.IsActive,
	})
	if err != nil {
		return result, err
	}

	result = Todo(res)

	return result, nil
}

func (p todoService) UpdateTodo(ctx context.Context, req UpdateTodo) (Todo, error) {
	var result = Todo{}

	res, err := p.tr.UpdateTodo(ctx, domain.UpdateTodo{
		ID:       req.ID,
		Title:    req.Title,
		Priority: req.Priority,
		IsActive: req.IsActive,
	})
	if err != nil {
		return result, err
	}

	result = Todo(res)

	return result, nil
}

func (p todoService) DeleteTodo(ctx context.Context, req DeleteTodo) (Todo, error) {
	var result = Todo{}

	res, err := p.tr.DeleteTodo(ctx, domain.DeleteTodo{
		ID: req.ID,
	})
	if err != nil {
		return result, err
	}

	result = Todo(res)

	return result, nil
}
