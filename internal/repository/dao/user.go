package dao

import (
	"context"

	"github.com/GoSimplicity/code_huihui/internal/model"
	"github.com/samber/do"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserDAO interface {
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
}

type userDAO struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserDAO(i *do.Injector) (UserDAO, error) {
	return &userDAO{
		db:     do.MustInvoke[*gorm.DB](i),
		logger: do.MustInvoke[*zap.Logger](i),
	}, nil
}

func (d *userDAO) FindByID(ctx context.Context, id int64) (*model.User, error) {
	return nil, nil
}

func (d *userDAO) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (d *userDAO) Create(ctx context.Context, user *model.User) error {
	return nil
}

func (d *userDAO) Update(ctx context.Context, user *model.User) error {
	return nil
}

func (d *userDAO) Delete(ctx context.Context, id int64) error {
	return nil
}
