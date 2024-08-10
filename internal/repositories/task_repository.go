package repositories

import (
	"context"
	"task-tracker/internal/db"
	"task-tracker/internal/models"
	"time"
)

type TaskRepository interface {
	AddTask(*models.Task) error
	AddTasksFromTemplate(*models.Template) error
	UpdateTask(*models.Task) error
	DeleteTask(id int) error
	GetTasks() ([]models.Task, error)
}

type taskRepository struct {
	psqlClient db.PsqlClient
}

func NewTaskRepository(dbcli db.PsqlClient) TaskRepository {
	return &taskRepository{
		psqlClient: dbcli,
	}
}

func (t *taskRepository) AddTask(task *models.Task) error {
	query := "INSERT INTO task (description, done, date) values ($1, $2, $3)"
	ctx := context.Background()
	if _, err := t.psqlClient.Db.Exec(ctx, query, task.Description, false, time.Now()); err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) AddTasksFromTemplate(tmpl *models.Template) error {
	ctx := context.Background()
	tx, err := t.psqlClient.Db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	query := "INSERT INTO task (description, done, date) values($1, $2, $3, $4)"

	for _, task := range tmpl.Tasks {
		if _, err := t.psqlClient.Db.Exec(ctx, query, task.Description, false, time.Now().Unix()); err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) UpdateTask(task *models.Task) error {
	query := "UPDATE task SET description = $1, done = $2 WHERE id = $3"
	ctx := context.Background()
	if _, err := t.psqlClient.Db.Exec(ctx, query, task.Description, task.Done, task.ID); err != nil {
		return err
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
