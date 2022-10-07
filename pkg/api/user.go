package api

import (
	"errors"
	"strings"
)

type UserService interface {
	New(user NewUserRequest) error
}

type UserRepository interface {
	CreateUser(NewUserRequest) error
}

type userService struct {
	storage UserRepository
}

func (u *userService) New(user NewUserRequest) error {
	// do some basic validations
	if user.Email == "" {
		return errors.New("user service - email required")
	}

	if user.Name == "" {
		return errors.New("user service - name required")
	}

	if user.WeightGoal == "" {
		return errors.New("user service - weight goal required")
	}

	// do some basic normalisation
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	err := u.storage.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		storage: userRepo,
	}
}
