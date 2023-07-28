package authentication

import (
	"context"
	"echo-modarch/database"
	"log"
)

type UserRepository struct {
	database.Repository
	conn *database.PostgresDatabase
}

func NewUserRepository() *UserRepository {
	conn := database.GetDatabaseConnection()

	return &UserRepository{
		conn: conn,
	}
}

func (this *UserRepository) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	rows, err := this.GetDB().QueryContext(ctx, "SELECT id, name, email, password FROM users WHERE email = $1", email)

	if err != nil {
		return nil, err
	}

	defer func() { // Alows to validate the error after the function returns
		err := rows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = User{}
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err == nil {
			return &user, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) InsertUser(ctx context.Context, user *User) error {
	_, err := repo.GetDB().ExecContext(
		ctx,
		"INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)",
		user.Id, user.Name, user.Email, user.Password,
	)
	return err
}
