package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/meis1kqt/fastapiprac/internal/models"
)


type TaskStore struct {
	db *sqlx.DB
}

func NewTaskStore(db *sqlx.DB) *TaskStore  {

	return &TaskStore{db:db}

}


func (t *TaskStore) GetAll() ([]models.Task, error) {

	var tasks []models.Task

	err := t.db.Select(&tasks, "SELECT * FROM tasks")

	if err != nil {
		return nil, err
	}

	return tasks, nil

}


func (t *TaskStore) GetById(id int) (*models.Task, error) {
	var task models.Task

	err := t.db.Get(&task, "SELECT * FROM tasks WHERE id = $1", id)

	if err == sql.ErrNoRows {
		return nil , nil
	}

	if err != nil {
		 return nil, err
	}

	return &task, nil
}

func (t *TaskStore) Create(input models.CreateTaskInput) error {



	query := `INSERT INTO tasks (title, description, completed, created_at) VAlUES($1,$2,$3,$4)`

	now := time.Now()

	err := t.db.QueryRowx(query, input.Title, input.Description, input.Completed, now)

	if err != nil {
		return fmt.Errorf("bag")
	}
	return nil
}

func (t *TaskStore) Delete(id int) error {

	query := `DELETE from tasks where id = $1`

	_ , err := t.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil




}