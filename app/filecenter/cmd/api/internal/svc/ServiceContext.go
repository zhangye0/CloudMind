package svc

import (
	"CloudMind/app/filecenter/cmd/api/internal/config"
	"CloudMind/app/filecenter/cmd/rpc/filecenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	FileRpc filecenter.Filecenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c,
		filecenter.NewFilecenter(zrpc.MustNewClient(c.FileRpcConf)),
	}
}
