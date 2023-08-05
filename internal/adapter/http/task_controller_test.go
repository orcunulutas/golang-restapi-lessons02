// go-restapi/internal/adapter/http/task_controller_test.go

package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	apphttp "go-restapi/internal/adapter/http"
	"go-restapi/internal/domain/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskController_GetAllTasks(t *testing.T) {
	taskUsecase := &mockTaskUsecase{}
	taskHandler := apphttp.NewTaskHandler(taskUsecase)

	// Simulate the request
	req, err := http.NewRequest("GET", "/tasks", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request and record the response
	taskHandler.GetAllTasks(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var tasks []model.Task
	err = json.Unmarshal(rr.Body.Bytes(), &tasks)
	assert.NoError(t, err)
	assert.Len(t, tasks, 3) // Assuming we have 2 tasks in the mockTaskUsecase
}

func TestTaskController_CreateTask(t *testing.T) {
	taskUsecase := &mockTaskUsecase{}
	taskHandler := apphttp.NewTaskHandler(taskUsecase)

	// Sample task data for creating a task
	taskData := model.Task{
		Title:       "New Task",
		Description: "New Task Description",
	}

	// Convert the taskData to JSON
	body, err := json.Marshal(taskData)
	assert.NoError(t, err)

	// Simulate the request with JSON body
	req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request and record the response
	taskHandler.CreateTask(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusCreated, rr.Code)
}

// You can similarly add unit tests for other methods of the TaskController.

// Mock TaskUsecase for testing
type mockTaskUsecase struct {
	// Burada diğer alanlar yer alabilir
	tasks []model.Task
}

func (m *mockTaskUsecase) GetAllTasks() ([]model.Task, error) {
	return []model.Task{
		{ID: 1, Title: "task1", Description: "market", CreatedAt: time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC)},
		{ID: 2, Title: "task2", Description: "kitapOku", CreatedAt: time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC)},
		{ID: 3, Title: "task3", Description: "tadilat", CreatedAt: time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(2022, 9, 23, 0, 0, 0, 0, time.UTC)},
	}, nil
}

func (m *mockTaskUsecase) CreateTask(task *model.Task) error {
	m.tasks = append(m.tasks, *task)
	return nil
}

func (m *mockTaskUsecase) DeleteTask(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			// Task'ı listeden çıkar
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("Task not found")
}
func (m *mockTaskUsecase) GetTaskByID(id int) (*model.Task, error) {
	for _, task := range m.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("Task not found")
}

func (m *mockTaskUsecase) UpdateTask(task *model.Task) error {
	for i, t := range m.tasks {
		if t.ID == task.ID {
			m.tasks[i] = *task
			return nil
		}
	}
	return errors.New("Task not found")
}

// Implement other methods of TaskUsecase similarly.
