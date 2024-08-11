package services

import (
	"task-tracker/internal/dto"
	"task-tracker/internal/models"
	"task-tracker/internal/repositories"
)

type TemplateService interface {
	AddTemplate(*dto.AddTmplReq) error
	UpdateTemplate(*dto.UpdateTmplReq) error
	DeleteTemplate(*dto.DeleteTmplReq) error
	GetTemplates() ([]models.Template, error)
}

type templateService struct {
	repo repositories.TemplateRepository
}

func NewTemplateRepository(repo repositories.TemplateRepository) TemplateService {
	return &templateService{
		repo: repo,
	}
}

func (t *templateService) AddTemplate(req *dto.AddTmplReq) error {
	tasks := make([]models.TemplateTask, len(req.Tasks))

	for i, task := range req.Tasks {
		tasks[i] = models.TemplateTask{
			TemplateID:  task.TemplateID,
			Description: task.Description,
		}
	}

	template := &models.Template{
		Name:  req.Name,
		Tasks: tasks,
	}

	return t.repo.AddTemplate(template)
}

func (t *templateService) UpdateTemplate(req *dto.UpdateTmplReq) error {
	tasks := make([]models.TemplateTask, len(req.Tasks))

	for i, task := range req.Tasks {
		tasks[i] = models.TemplateTask{
			TemplateID:  task.TemplateID,
			Description: task.Description,
		}
	}

	template := &models.Template{
		Name:  req.Name,
		Tasks: tasks,
	}

	return t.repo.UpdateTemplate(template)
}

func (t *templateService) DeleteTemplate(req *dto.DeleteTmplReq) error {
	return t.repo.DeleteTemplate(req.ID)
}

func (t *templateService) GetTemplates() ([]models.Template, error) {
	return t.repo.GetTemplates()
}
