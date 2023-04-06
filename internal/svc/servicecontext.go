package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"zShop/internal/config"
	"zShop/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
	}
}
