package database

import (
	"database/sql"
	"echo-modarch/utils"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	db *sql.DB
}

func NewPostgresDatabase() (*PostgresDatabase, error) {
	config := utils.GetConfig()

	urlDB := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.PostgresDB,
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
