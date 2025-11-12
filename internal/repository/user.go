package repository

import (
	"context"

	"github.com/GoSimplicity/code_huihui/internal/model"
	"github.com/GoSimplicity/code_huihui/internal/repository/dao"
	"github.com/samber/do"
	"go.uber.org/zap"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
}

type userRepository struct {
	dao    dao.UserDAO
	logger *zap.Logger
}

func NewUserRepository(i *do.Injector) (UserRepository, error) {
	userDAO := do.MustInvoke[dao.UserDAO](i)
	logger := do.MustInvoke[*zap.Logger](i)
	return &userRepository{
		dao:    userDAO,
		logger: logger,
	}, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	return nil, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	return nil
}
