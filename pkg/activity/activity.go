package activity

import (
	"context"
	"database/sql"

	"github.com/gadhittana01/todolist/config"
	"github.com/gadhittana01/todolist/pkg/domain"
)

type IResource interface {
	GetActivities(ctx context.Context) ([]domain.Activity, error)
	GetActivityByID(ctx context.Context, req domain.GetActivityByID) (domain.Activity, error)
	CreateActivity(ctx context.Context, req domain.CreateActivity) (domain.Activity, error)
	UpdateActivity(ctx context.Context, req domain.UpdateActivity) (domain.Activity, error)
	DeleteActivity(ctx context.Context, req domain.DeleteActivity) (domain.Activity, error)
}

type module struct {
	persistent persistent
}

func New(cfg *config.GlobalConfig, db *sql.DB) IResource {
	return &module{
		persistent: newPersistent(db),
	}
}

func (m module) GetActivities(ctx context.Context) ([]domain.Activity, error) {
	return m.persistent.getActivities(ctx)
}

func (m module) GetActivityByID(ctx context.Context, req domain.GetActivityByID) (domain.Activity, error) {
	return m.persistent.getActivityByID(ctx, req)
}

func (m module) CreateActivity(ctx context.Context, req domain.CreateActivity) (domain.Activity, error) {
	return m.persistent.createActivity(ctx, req)
}

func (m module) UpdateActivity(ctx context.Context, req domain.UpdateActivity) (domain.Activity, error) {
	return m.persistent.updateActivity(ctx, req)
}

func (m module) DeleteActivity(ctx context.Context, req domain.DeleteActivity) (domain.Activity, error) {
	return m.persistent.deleteActivity(ctx, req)
}
