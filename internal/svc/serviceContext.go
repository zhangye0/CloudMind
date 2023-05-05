package svc

import (
	"CloudMind/internal/config"
	"CloudMind/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserInfoModel model.UserinfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		UserInfoModel: model.NewUserinfoModel(sqlConn, c.Cache),
	}
}
