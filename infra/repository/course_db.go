package repository

import (
	"database/sql"
	"imersao-dev-full-cycle-go/entity"
)

type CourseDatabaseRepository struct {
	Db *sql.DB
}

func (c CourseDatabaseRepository) Insert(course entity.Course) error {
	statement, err := c.Db.Prepare(`Insert into courses(id, name, description, status) values(?,?,?,?)`)

	if err != nil {
		return err
	}

	_, err = statement.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	if err != nil {
		return err
	}
	return nil
}
