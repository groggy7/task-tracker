package handlers

import (
	"log"
	"net/http"
	"task-tracker/internal/dto"
	"task-tracker/internal/services"

	"github.com/gin-gonic/gin"
)

type TemplateHandler interface {
	AddTemplate(*gin.Context)
	UpdateTemplate(*gin.Context)
	DeleteTemplate(*gin.Context)
	GetTemplates(*gin.Context)
}

type templateHandler struct {
	templateService services.TemplateService
}

func NewTemplateHandler(templateService services.TemplateService) TemplateHandler {
	return &templateHandler{
		templateService: templateService,
	}
}

func (h *templateHandler) AddTemplate(ctx *gin.Context) {
	req := new(dto.AddTmplReq)
	if err := ctx.BindJSON(req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		return
	}

	if err := h.templateService.AddTemplate(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *templateHandler) UpdateTemplate(ctx *gin.Context) {
	req := new(dto.UpdateTmplReq)
	if err := ctx.BindJSON(req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		return
	}

	if err := h.templateService.UpdateTemplate(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *templateHandler) DeleteTemplate(ctx *gin.Context) {
	req := new(dto.DeleteTmplReq)
	if err := ctx.BindJSON(req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		return
	}

	if err := h.templateService.DeleteTemplate(req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *templateHandler) GetTemplates(ctx *gin.Context) {
	templates, err := h.templateService.GetTemplates()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, templates)
}
