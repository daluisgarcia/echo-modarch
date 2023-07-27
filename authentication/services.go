package authentication

import (
	"context"
	"echo-modarch/utils"
	"fmt"

	"github.com/google/uuid"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: NewUserRepository(),
	}
}

func (this *UserService) LoginUser(ctx context.Context, user *LoginUserRequest) (*User, error) {
	userInDB, err := this.repo.FindUserByEmail(ctx, user.Email)

	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	err = utils.CompareHashedString(userInDB.Password, user.Password)

	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")
	}

	return userInDB, nil
}

func (this *UserService) RegisterUser(ctx context.Context, user *RegisterUserRequest) (*User, error) {

	hashedPassword, err := utils.HashString(user.Password)

	if err != nil {
		return nil, err
	}

	userToInsert := &User{
		Id:       uuid.New().String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	err = this.repo.InsertUser(ctx, userToInsert)

	if err != nil {
		return nil, fmt.Errorf("Unexpected error creating user")
	}

	return userToInsert, nil
}
