package svc

import (
	"CloudMind/app/usercenter/cmd/rpc/internal/config"
	"CloudMind/app/usercenter/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Redis
	GormDB        *gorm.DB
	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDB, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀
			SingularTable: true, // 使用单数表
		},
	})
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
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserAuthModel: model.NewUserAuthModel(gormDB),
		UserModel:     model.NewUserModel(gormDB),
	}
}
