package repositories

import (
	"context"
	"task-tracker/internal/db"
	"task-tracker/internal/models"
)

type TemplateRepository interface {
	AddTemplate(tmpl *models.Template) error
	UpdateTemplate(tmpl *models.Template) error
	GetTemplates() ([]models.Template, error)
}

type templateRepository struct {
	psqlClient db.PsqlClient
}

func NewTemplateRepository(dbcli db.PsqlClient) TemplateRepository {
	return &templateRepository{
		psqlClient: dbcli,
	}
}

func (t *templateRepository) AddTemplate(tmpl *models.Template) error {
	ctx := context.Background()
	tx, err := t.psqlClient.Db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	var templateID int
	query := "INSERT INTO template (name) values($1) RETURNING id"
	if err := tx.QueryRow(ctx, query, tmpl.Name).Scan(&templateID); err != nil {
		return err
	}

	for _, task := range tmpl.Tasks {
		query := "INSERT INTO template_task (template_id, description) values($1, $2)"
		if _, err := tx.Exec(ctx, query, task.TemplateID, task.Description); err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (t *templateRepository) UpdateTemplate(tmpl *models.Template) error {
	ctx := context.Background()
	tx, err := t.psqlClient.Db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	query := "UPDATE template SET name = $1"
	if _, err := tx.Exec(ctx, query, tmpl.Name); err != nil {
		return err
	}

	for _, task := range tmpl.Tasks {
		query := "UPDATE template_task SET description = $1"
		if _, err := tx.Exec(ctx, query, task.Description); err != nil {
			return err
		}
	}

	return nil
}

func (t *templateRepository) GetTemplates() ([]models.Template, error) {
	ctx := context.Background()
	query := "SELECT * FROM template"
	rows, err := t.psqlClient.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	templates := make([]models.Template, 0)
	for rows.Next() {
		template := models.Template{}
		if err := rows.Scan(&template.ID, &template.Name); err != nil {
			return nil, err
		}

		query = "SELECT * FROM task_template WHERE template_id = $1"
		task_rows, err := t.psqlClient.Db.Query(ctx, query, template.ID)
		if err != nil {
			return nil, err
		}

		tasks := make([]models.TemplateTask, 0)
		for task_rows.Next() {
			task := models.TemplateTask{}
			if err := task_rows.Scan(&task.ID, &task.TemplateID, &task.Description); err != nil {
				return nil, err
			}
			tasks = append(tasks, task)
		}
		template.Tasks = tasks
		templates = append(templates, template)
	}

	return templates, nil
}
