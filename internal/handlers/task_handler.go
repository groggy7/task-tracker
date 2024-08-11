package handlers

import (
	"log"
	"net/http"
	"task-tracker/internal/dto"
	"task-tracker/internal/services"

	"github.com/gin-gonic/gin"
)

type TaskHandler interface {
	AddTask(*gin.Context)
	AddTasksFromTemplate(*gin.Context)
	UpdateTask(*gin.Context)
	DeleteTask(*gin.Context)
	GetTasks(*gin.Context)
}

type taskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) TaskHandler {
	return &taskHandler{
		taskService: taskService,
	}
}

func (h *taskHandler) AddTask(ctx *gin.Context) {
	req := new(dto.AddTaskReq)
	if err := ctx.BindJSON(req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		return
	}

	if err := h.taskService.AddTask(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *taskHandler) AddTasksFromTemplate(ctx *gin.Context) {
	req := new(dto.AddFromTmplReq)
	if err := ctx.BindJSON(req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		return
	}

	if err := h.taskService.AddTasksFromTemplate(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *taskHandler) UpdateTask(ctx *gin.Context) {
	req := new(dto.UpdateTaskReq)
	if err := ctx.BindJSON(req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		return
	}

	if err := h.taskService.UpdateTask(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *taskHandler) DeleteTask(ctx *gin.Context) {
	req := new(dto.DeleteTaskReq)
	if err := ctx.BindJSON(req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		return
	}

	if err := h.taskService.DeleteTask(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *taskHandler) GetTasks(ctx *gin.Context) {
	tasks, err := h.taskService.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
