package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := "INSERT INTO todos ( content, user_id, created_at) VALUES ($1, $2, $3)"
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := "SELECT * FROM todos WHERE id = $1"
	err = Db.QueryRow(cmd, id).Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	todos = []Todo{}
	cmd := "SELECT * FROM todos"

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)

		}
		todos = append(todos, todo)
	}
	return todos, err
}

func (u *User) GetTodoByUser() (todos []Todo, err error) {
	todos = []Todo{}
	cmd := "SELECT * FROM todos WHERE user_id = $1"

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)

		}
		todos = append(todos, todo)
	}
	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := "UPDATE todos set content = $1, user_id = $2 WHERE id = $3"
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := "DELETE from todos WHERE id = $1"
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
