package auth

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Register(login, password, confirm string) error {
	if login == "" || password == "" {
		return errors.New("логин и пароль обязательны")
	}

	if password != confirm {
		return errors.New("пароли не совпадают")
	}

	if _, err := s.repo.GetByLogin(login); err == nil {
		return errors.New("пользователь уже существует")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error:", err)
		return err
	}

	return s.repo.CreateUser(login, string(hash))
}

func (s *Service) Login(login, password string) (*User, error) {
	user, err := s.repo.GetByLogin(login)
	if err != nil {
		return nil, errors.New("неверный логин или пароль")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("неверный логин или пароль")
	}

	return user, nil
}
