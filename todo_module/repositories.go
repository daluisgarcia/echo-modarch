package todo_module

import (
	"context"
	"echo-modarch/database"
	"log"
)

type ToDoRepository struct {
	database.Repository
	conn *database.PostgresDatabase
}

func NewToDoRepository() *ToDoRepository {
	conn := database.GetDatabaseConnection()

	return &ToDoRepository{
		conn: conn,
	}
}

func (this *ToDoRepository) GetAllTasks(ctx context.Context, user_id string) ([]*Task, error) {
	rows, err := this.GetDB().QueryContext(ctx, "SELECT id, title, description, is_done FROM tasks where user_id = $1", user_id)

	if err != nil {
		return nil, err
	}

	defer func() { // Alows to validate the error after the function returns
		err := rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	toDosList := []*Task{}
	for rows.Next() {
		task := &Task{}
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Done); err == nil {
			toDosList = append(toDosList, task)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return toDosList, nil
}

func (this *ToDoRepository) GetTaskById(ctx context.Context, id string) (*Task, error) {
	rows, err := this.GetDB().QueryContext(ctx, "SELECT id, title, description, is_done FROM tasks WHERE id = $1", id)

	if err != nil {
		return nil, err
	}

	defer func() { // Alows to validate the error after the function returns
		err := rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		task := &Task{}
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Done); err != nil {
			return nil, err
		}
		return task, nil
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nil, err
}

func (this *ToDoRepository) InsertTask(ctx context.Context, task Task, userId string) error {
	_, err := this.GetDB().ExecContext(
		ctx,
		"INSERT INTO tasks (id, title, description, user_id) VALUES ($1, $2, $3, $4)",
		task.Id, task.Title, task.Description, userId,
	)

	return err
}
