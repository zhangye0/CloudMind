package svc

import (
	"CloudMind/app/es/cmd/rpc/es"
	"CloudMind/app/usercenter/cmd/rpc/internal/config"
	"CloudMind/app/usercenter/model"
	"CloudMind/common/gormlogger"
	"github.com/geiqin/thirdparty/oauth"
	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type ServiceContext struct {
	Config          config.Config
	Redis           *redis.Redis
	GormDB          *gorm.DB
	Cache           *collection.Cache
	Bloom           *bloom.Filter
	wxAuth          *oauth.AuthWxWechat
	UserModel       model.UserModel
	UserAuthModel   model.UserAuthModel
	UserAvatarModel model.UserAvatarModel
	EsRpc           es.Es
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDB, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀
			SingularTable: true, // 使用单数表
		},
		Logger: gormlogger.New(gormlogger.Config{LogLevel: logger.Info}),
	})

	wxConf := &oauth.AuthConfig{
		ClientId:     "wx4f9b6ad3ffd3a295",
		ClientSecret: "a1df664c4379e1a71aff7af68aed3c25",
		RedirectUrl:  "",
	}

	if err != nil {
		panic(err)
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(64)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(64)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxIdleTime(time.Minute)
	// SetConnMaxLifetime 设置了连接存活的最大时间。
	db.SetConnMaxLifetime(time.Minute)

	cache, err := collection.NewCache(time.Minute, collection.WithLimit(10000))
	Redis := redis.MustNewRedis(c.RedisConf)

	return &ServiceContext{
		Config:          c,
		Redis:           Redis,
		Cache:           cache,
		Bloom:           bloom.New(Redis, "bloom", 1024),
		wxAuth:          oauth.NewAuthWxWechat(wxConf),
		UserAuthModel:   model.NewUserAuthModel(gormDB),
		UserModel:       model.NewUserModel(gormDB),
		UserAvatarModel: model.NewUserAvatarModel(gormDB),
		EsRpc:           es.NewEs(zrpc.MustNewClient(c.EsRpcConf)),
	}
}
