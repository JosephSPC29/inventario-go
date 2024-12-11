package services

import (
	"errors"
	"inventario-go/models"
	"inventario-go/repositories"
	"inventario-go/utils"
)

type UserService interface {
	RegisterUser(username, password, email, role string) error
	AuthenticateUser(username, password string) (string, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) RegisterUser(username, password, email, role string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Usuario:    username,
		Contraseña: hashedPassword,
		Email:      email,
		Rol:        role,
	}

	return s.repo.CreateUser(user)
}

func (s *userService) AuthenticateUser(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(password, user.Contraseña) {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Rol)
	if err != nil {
		return "", err
	}

	return token, nil
}
