package storage

import (
	"database/sql"
	"fmt"

	"github.com/NorousS/CRUD/internal/models"
)

func GetAllTodo() ([]models.Todo, error) {
	rows, err := models.DB.Query(`SELECT id, title, description, completed FROM todos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, err
}

func GetTodoByID(id int) (*models.Todo, error) {
	var todo models.Todo
	err := models.DB.QueryRow(`SELECT id, title, description, completed FROM todos WHERE id = $1`, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func CreateTodo(todo *models.Todo) error {
	query := `INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id`
	err := models.DB.QueryRow(query, todo.Title, todo.Description, todo.Completed).Scan(&todo.ID)
	return err
}

func UpdateTodo(id int, todo *models.Todo) error {
	var updatedID int
	query := `UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4 RETURNING id`
	err := models.DB.QueryRow(query, todo.Title, todo.Description, todo.Completed, id).Scan(&updatedID)

	if err == sql.ErrNoRows {
		return fmt.Errorf("todo with id %d not found", id)
	}
	return err
}

func DeleteTodo(id int) error {
	var deletedID int
	err := models.DB.QueryRow(`DELETE FROM todos WHERE id = $1 RETURNING id`, id).Scan(&deletedID)
	if err == sql.ErrNoRows {
		return fmt.Errorf("todo with id %d not found", id)
	}
	return err
}
