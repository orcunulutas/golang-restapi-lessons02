package http

import (
	"encoding/json"
	"fmt"
	"go-restapi/internal/domain/model"
	"go-restapi/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type TaskHandler interface {
	GetAllTasks(w http.ResponseWriter, r *http.Request)
	GetTaskByID(w http.ResponseWriter, r *http.Request)
	CreateTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase}
}

func (h *taskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskUsecase.GetAllTasks()
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, tasks)
}

func (h *taskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id := parseID(idStr)
	if id == 0 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.taskUsecase.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (h *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	err = h.taskUsecase.CreateTask(&task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *taskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id := parseID(idStr)
	if id == 0 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	task.ID = id
	err = h.taskUsecase.UpdateTask(&task)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *taskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id := parseID(idStr)
	if id == 0 {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err := h.taskUsecase.DeleteTask(id)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parseID(idStr string) int {
	id := 0
	fmt.Sscanf(idStr, "%d", &id)
	return id
}

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
