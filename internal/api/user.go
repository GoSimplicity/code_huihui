package api

import (
	"github.com/GoSimplicity/code_huihui/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"go.uber.org/zap"
)

type UserHandler struct {
	svc    service.UserService
	logger *zap.Logger
}

func NewUserHandler(i *do.Injector) (*UserHandler, error) {
	svc := do.MustInvoke[service.UserService](i)
	logger := do.MustInvoke[*zap.Logger](i)
	return &UserHandler{
		svc:    svc,
		logger: logger,
	}, nil
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	userGroup := server.Group("/api/user")
	{
		userGroup.POST("/login", h.Login)
		userGroup.POST("/register", h.Register)
		userGroup.GET("/get", h.GetUser)
		userGroup.POST("/update", h.UpdateUser)
		userGroup.DELETE("/delete", h.DeleteUser)
	}
}

func (h *UserHandler) Login(c *gin.Context) {
}

func (h *UserHandler) GetUser(c *gin.Context) {
}

func (h *UserHandler) Register(c *gin.Context) {
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
}
