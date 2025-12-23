package users

import (
    "context"
    "errors"
)

type Service interface {
    GetProfile(ctx context.Context, id int64) (*User, error)
    UpdateProfile(ctx context.Context, u *User) error
}

type service struct {
    repo Repository
}

func NewService(repo Repository) Service {
    return &service{repo: repo}
}

func (s *service) GetProfile(ctx context.Context, id int64) (*User, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *service) UpdateProfile(ctx context.Context, u *User) error {
    if u.Age == 0 {
        return errors.New("age cannot be empty")
    }
    if u.Contact == "" {
        return errors.New("contact cannot be empty")
    }
    if u.PrimeTime == "" {
        return errors.New("prime time cannot be empty")
    }
    if len(u.Games) == 0 {
        return errors.New("games cannot be empty")
    }

    return s.repo.Update(ctx, u)
}
