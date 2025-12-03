package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(email, password, confirm string) error {
	if password != confirm {
		return errors.New("password mismatch")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(email, string(hash))
}

func (s *Service) Login(email, password string) (*User, error) {
	u, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return u, nil
}
