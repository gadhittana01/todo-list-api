package resthttp

import (
	"github.com/go-chi/chi"
)

type RouterDependencies struct {
	AS ActivityService
	TS TodoService
}

func NewRoutes(rd RouterDependencies) *chi.Mux {
	router := chi.NewRouter()

	ah := newActivityHandler(rd.AS)
	th := newTodoHandler(rd.TS)

	// activity
	router.Get("/activity-groups", ah.GetActivities)
	router.Get("/activity-groups/{id}", ah.GetActivityByID)
	router.Post("/activity-groups", ah.CreateActivity)
	router.Patch("/activity-groups/{id}", ah.UpdateActivity)
	router.Delete("/activity-groups/{id}", ah.DeleteActivity)

	// todo
	router.Get("/todo-items", th.GetTodos)
	router.Get("/todo-items/{id}", th.GetTodoByID)
	router.Post("/todo-items", th.CreateTodo)
	router.Patch("/todo-items/{id}", th.UpdateTodo)
	router.Delete("/todo-items/{id}", th.DeleteTodo)

	return router
}
