package svc

import (
	"CloudMind/app/fileservice/cmd/rpc/internal/config"
	"CloudMind/app/fileservice/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	FileModel model.FileModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:    c,
		FileModel: model.NewFileModel(sqlConn, c.Cache),
	}
}
