package svc

import "github.com/wuntsong-org/go-zero-plus/tools/goctlwt/example/rpc/hi/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
