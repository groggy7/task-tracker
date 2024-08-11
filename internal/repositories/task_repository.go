package repositories

import (
	"context"
	"fmt"
	"task-tracker/internal/db"
	"task-tracker/internal/models"
	"time"
)

type TaskRepository interface {
	AddTask(*models.Task) error
	AddTasksFromTemplate(int) error
	UpdateTask(*models.Task) error
	DeleteTask(id int) error
	GetTasks() ([]models.Task, error)
}

type taskRepository struct {
	psqlClient *db.PsqlClient
}

func NewTaskRepository(dbcli *db.PsqlClient) TaskRepository {
	return &taskRepository{
		psqlClient: dbcli,
	}
}

func (t *taskRepository) AddTask(task *models.Task) error {
	query := "INSERT INTO task (description, done, date) values ($1, $2, $3)"
	ctx := context.Background()
	if _, err := t.psqlClient.Db.Exec(ctx, query, task.Description, false, time.Now().Unix()); err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) AddTasksFromTemplate(templateID int) error {
	ctx := context.Background()
	query := "SELECT * FROM template WHERE id = $1"
	row := t.psqlClient.Db.QueryRow(ctx, query, templateID)

	template := models.Template{}
	if err := row.Scan(&template.ID, &template.Name); err != nil {
		return err
	}

	query = "SELECT * FROM template_task WHERE template_id = $1"
	rows, err := t.psqlClient.Db.Query(ctx, query, templateID)
	if err != nil {
		return err
	}

	templateTasks := make([]models.TemplateTask, 0)
	for rows.Next() {
		task := models.TemplateTask{}
		if err := rows.Scan(&task.ID, &task.TemplateID, &task.Description); err != nil {
			return err
		}
		templateTasks = append(templateTasks, task)
	}

	for _, task := range templateTasks {
		query = "INSERT INTO task (description, done, date) values ($1, $2, $3)"
		if _, err := t.psqlClient.Db.Exec(ctx, query, task.Description, false, time.Now().Unix()); err != nil {
			return err
		}
	}

	return nil
}

func (t *taskRepository) UpdateTask(task *models.Task) error {
	query := "UPDATE task SET description = $1, done = $2 WHERE id = $3"
	ctx := context.Background()
	result, err := t.psqlClient.Db.Exec(ctx, query, task.Description, task.Done, task.ID)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated, task ID might be incorrect: %d", task.ID)
	}

	return nil
}

func (t *taskRepository) DeleteTask(id int) error {
	query := "DELETE from task WHERE id = $1"
	ctx := context.Background()
	if _, err := t.psqlClient.Db.Query(ctx, query, id); err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetTasks() ([]models.Task, error) {
	query := "SELECT * FROM task"
	ctx := context.Background()
	rows, err := t.psqlClient.Db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	tasks := make([]models.Task, 0)
	for rows.Next() {
		task := models.Task{}
		if err := rows.Scan(&task.ID, &task.Description, &task.Done, &task.Date); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
