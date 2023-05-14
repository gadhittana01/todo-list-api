package main

import (
	"github.com/gadhittana01/todolist/config"
	"github.com/gadhittana01/todolist/db"
	"github.com/gadhittana01/todolist/handler/resthttp"
	"github.com/gadhittana01/todolist/pkg/activity"
	"github.com/gadhittana01/todolist/pkg/todo"
	"github.com/gadhittana01/todolist/services"
)

func initApp(c *config.GlobalConfig) error {
	db := db.InitDB()
	activityPkg := activity.New(c, db)
	todoPkg := todo.New(c, db)

	as, err := services.NewActivityService(services.ActivityDependencies{
		AR: activityPkg,
	})
	if err != nil {
		return err
	}

	ts, err := services.NewTodoService(services.TodoDependencies{
		TR: todoPkg,
	})
	if err != nil {
		return err
	}

	return startHTTPServer(resthttp.NewRoutes(resthttp.RouterDependencies{
		AS: as,
		TS: ts,
	}), c)
}
