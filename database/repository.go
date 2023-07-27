package database

import "database/sql"

type Repository struct{}

func (this *Repository) GetDB() *sql.DB {
	return GetDatabaseConnection().GetDB()
}
