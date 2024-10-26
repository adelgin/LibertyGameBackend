package service

import (
	"context"
	"fmt"
	"time"

	"libertyGame/internal/repository"
)

type UserService interface {
	GetUserByID(ctx context.Context, id int64) (*repository.User, error)
	AddUser(ctx context.Context, user *repository.User) error
	CountOfAllUsers(ctx context.Context) (int64, error)
	GetRefsOfUserFromID(ctx context.Context, id int64) ([]repository.User, error)
	CountRefsOfUserFromID(ctx context.Context, id int64) (int64, error)
	GetTopOfRefs(ctx context.Context, count int) ([]repository.User, error)
}

type userService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUserByID(ctx context.Context, id int64) (*repository.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *userService) AddUser(ctx context.Context, user *repository.User) error {
	user.CreatedAt = time.Now() // Set CreatedAt before adding
	return s.repo.AddUser(ctx, user)
}

func (s *userService) CountOfAllUsers(ctx context.Context) (int64, error) {
	return s.repo.CountOfAllUsers(ctx)
}

func (s *userService) GetRefsOfUserFromID(ctx context.Context, id int64) ([]repository.User, error) {
	return s.repo.GetRefsOfUserFromID(ctx, id)
}

func (s *userService) CountRefsOfUserFromID(ctx context.Context, id int64) (int64, error) {
	return s.repo.CountRefsOfUserFromID(ctx, id)
}

func (s *userService) GetTopOfRefs(ctx context.Context, count int) ([]repository.User, error) {
	if count <= 0 {
		return nil, fmt.Errorf("count must be a positive integer")
	}
	return s.repo.GetTopOfRefs(ctx, count)
}
