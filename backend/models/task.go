package models

import (
	"log"

	db "github.com/go-compose-rest/database"
)

type Task struct {
	ID          int64  `json:"id"`
	Description string `binding:"required" json:"description"`
	IsDone      *bool  `binding:"required" json:"isDone"`
}

func (task *Task) Save() error {
	createSql := `insert into tasks(description,is_done) values(?,?)`
	stmt, err := db.DB.Prepare(createSql)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(task.Description, task.IsDone)
	if err != nil {
		log.Fatal(err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	task.ID = id
	return nil
}

func GetTasks() ([]Task, error) {
	var tasks []Task
	querySql := "select * from tasks"
	rows, error := db.DB.Query(querySql)
	if error != nil {
		return nil, error
	}
	defer rows.Close()
	for rows.Next() {
		var task Task
		rows.Scan(&task.ID, &task.Description, &task.IsDone)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskById(id int64) (*Task, error) {
	querySql := `select * from tasks where id=?`
	row := db.DB.QueryRow(querySql, id)
	var task Task
	error := row.Scan(&task.ID, &task.Description, &task.IsDone)
	if error != nil {
		return nil, error
	}
	return &task, nil
}

func (task Task) UpdateTask() error {
	updateDql := `update tasks set is_done=? where id=?`
	stmt, error := db.DB.Prepare(updateDql)
	if error != nil {
		return error
	}
	defer stmt.Close()
	_, error = stmt.Exec(task.IsDone, task.ID)
	return error

}

func (task Task) DeleteTask() error {
	deleteSql := `delete from tasks where id=?`
	stmt, error := db.DB.Prepare(deleteSql)
	if error != nil {
		return error
	}
	defer stmt.Close()
	_, error = stmt.Exec(task.ID)
	return error
}
