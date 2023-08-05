package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	InitRoutes() http.Handler
}

type router struct {
	taskHandler TaskHandler
	userHandler UserHandler
}

func NewRouter(taskHandler TaskHandler, userHandler UserHandler) http.Handler {
	var r = router{
		taskHandler: taskHandler,
		userHandler: userHandler,
	}
	return r.InitRoutes()
}

func (r *router) InitRoutes() http.Handler {
	router := chi.NewRouter()

	router.Get("/tasks", r.taskHandler.GetAllTasks)
	router.Get("/tasks/{id}", r.taskHandler.GetTaskByID)
	router.Post("/tasks", r.taskHandler.CreateTask)
	router.Put("/tasks/{id}", r.taskHandler.UpdateTask)
	router.Delete("/tasks/{id}", r.taskHandler.DeleteTask)

	// Define user routes here

	return router
}
