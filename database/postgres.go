package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	db *sql.DB
}

func NewPostgresDatabase(postgresUser, postgresPassword, postgresHost, postgresDB string) (*PostgresDatabase, error) {

	urlDB := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		postgresUser,
		postgresPassword,
		postgresHost,
		postgresDB,
	)

	db, err := sql.Open("postgres", urlDB)

	if err != nil {
		return nil, err
	}

	return &PostgresDatabase{db: db}, nil
}

func (repo *PostgresDatabase) GetDB() *sql.DB {
	return repo.db
}

func (this *PostgresDatabase) Close() error {
	return this.db.Close()
}

var databaseConnection *PostgresDatabase

func SetDatabaseConnection(conn *PostgresDatabase) {
	databaseConnection = conn

	// Initialize tables
	query, err := ioutil.ReadFile("database/up.sql")

	if err != nil {
		panic(err)
	}
	if _, err := conn.db.Exec(string(query)); err != nil {
		panic(err)
	}
}

func GetDatabaseConnection() *PostgresDatabase {
	return databaseConnection
}
