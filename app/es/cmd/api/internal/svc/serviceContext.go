package svc

import (
	"CloudMind/app/es/cmd/api/internal/config"
	"CloudMind/app/es/cmd/rpc/es"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	EsRpc  es.Es
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		EsRpc:  es.NewEs(zrpc.MustNewClient(c.EsRpcConf)),
	}
}
