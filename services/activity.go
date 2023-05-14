package services

import (
	"context"

	"github.com/gadhittana01/todolist/pkg/domain"
)

type ActivityService interface {
	GetActivities(ctx context.Context) ([]Activity, error)
	GetActivityByID(ctx context.Context, req GetActivityByID) (Activity, error)
	CreateActivity(ctx context.Context, req CreateActivity) (Activity, error)
	UpdateActivity(ctx context.Context, req UpdateActivity) (Activity, error)
	DeleteActivity(ctx context.Context, req DeleteActivity) (Activity, error)
}

type activityService struct {
	ar ActivityResource
}

func NewActivityService(dep ActivityDependencies) (ActivityService, error) {
	return &activityService{
		ar: dep.AR,
	}, nil
}

func (p activityService) GetActivities(ctx context.Context) ([]Activity, error) {
	var result = []Activity{}

	res, err := p.ar.GetActivities(ctx)
	if err != nil {
		return result, err
	}

	for _, item := range res {
		result = append(result, Activity{
			ID:        item.ID,
			Title:     item.Title,
			Email:     item.Email,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return result, nil
}

func (p activityService) GetActivityByID(ctx context.Context, req GetActivityByID) (Activity, error) {
	var result = Activity{}

	res, err := p.ar.GetActivityByID(ctx, domain.GetActivityByID{
		ID: req.ID,
	})
	if err != nil {
		return result, err
	}

	result = Activity(res)

	return result, nil
}

func (p activityService) CreateActivity(ctx context.Context, req CreateActivity) (Activity, error) {
	var result = Activity{}

	res, err := p.ar.CreateActivity(ctx, domain.CreateActivity{
		Title: req.Title,
		Email: req.Email,
	})
	if err != nil {
		return result, err
	}

	result = Activity(res)

	return result, nil
}

func (p activityService) UpdateActivity(ctx context.Context, req UpdateActivity) (Activity, error) {
	var result = Activity{}

	res, err := p.ar.UpdateActivity(ctx, domain.UpdateActivity{
		ID:    req.ID,
		Title: req.Title,
	})
	if err != nil {
		return result, err
	}

	result = Activity(res)

	return result, nil
}

func (p activityService) DeleteActivity(ctx context.Context, req DeleteActivity) (Activity, error) {
	var result = Activity{}

	res, err := p.ar.DeleteActivity(ctx, domain.DeleteActivity{
		ID: req.ID,
	})
	if err != nil {
		return result, err
	}

	result = Activity(res)

	return result, nil
}
