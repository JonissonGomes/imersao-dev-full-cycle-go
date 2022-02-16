package repository

import (
	"database/sql"
)

type CourseDatabaseRepository struct {
	Db *sql.DB
}
