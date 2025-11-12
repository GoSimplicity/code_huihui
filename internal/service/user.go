package service

import (
	"context"

	"github.com/GoSimplicity/code_huihui/internal/model"
	"github.com/GoSimplicity/code_huihui/internal/repository"
	"github.com/samber/do"
	"go.uber.org/zap"
)

type UserService interface {
	Login(ctx context.Context, username, password string) (*model.User, error)
	GetUserByID(ctx context.Context, id int64) (*model.User, error)
	Register(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int64) error
}

type userService struct {
	repo   repository.UserRepository
	logger *zap.Logger
}

func NewUserService(i *do.Injector) (UserService, error) {
	repo := do.MustInvoke[repository.UserRepository](i)
	logger := do.MustInvoke[*zap.Logger](i)
	return &userService{
		repo:   repo,
		logger: logger,
	}, nil
}

func (s *userService) Login(ctx context.Context, username, password string) (*model.User, error) {
	return nil, nil
}

func (s *userService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	return nil, nil
}

func (s *userService) Register(ctx context.Context, user *model.User) error {
	return nil
}

func (s *userService) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return nil
}
