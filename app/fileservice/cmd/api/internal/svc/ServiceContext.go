package svc

import (
	"CloudMind/app/fileservice/cmd/api/internal/config"
	"CloudMind/app/fileservice/cmd/rpc/filecenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	FileRpc filecenter.Filecenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		FileRpc: filecenter.NewFilecenter(zrpc.MustNewClient(c.FileRpcConf)),
	}
}
