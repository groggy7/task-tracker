package services

import (
	"task-tracker/internal/dto"
	"task-tracker/internal/models"
	"task-tracker/internal/repositories"
)

type TaskService interface {
	AddTask(*dto.AddTaskReq) error
	AddTasksFromTemplate(*dto.AddFromTmplReq) error
	UpdateTask(*dto.UpdateTaskReq) error
	DeleteTask(*dto.DeleteTaskReq) error
	GetTasks() ([]models.Task, error)
}

type taskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (t *taskService) AddTask(req *dto.AddTaskReq) error {
	task := models.Task{
		Description: req.Description,
	}

	return t.repo.AddTask(&task)
}

func (t *taskService) AddTasksFromTemplate(req *dto.AddFromTmplReq) error {
	return t.repo.AddTasksFromTemplate(req.TemplateID)
}

func (t *taskService) UpdateTask(req *dto.UpdateTaskReq) error {
	updatedTask := &models.Task{
		ID:          req.ID,
		Description: req.Description,
		Done:        req.Done,
	}

	return t.repo.UpdateTask(updatedTask)
}

func (t *taskService) DeleteTask(req *dto.DeleteTaskReq) error {
	return t.repo.DeleteTask(req.ID)
}

func (t *taskService) GetTasks() ([]models.Task, error) {
	return t.repo.GetTasks()
}
