package svc

import (
	"CloudMind/app/filecenter/cmd/rpc/internal/config"
	"CloudMind/app/filecenter/model"
	"CloudMind/common/gormlogger"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	GormDB      *gorm.DB
	Cache       *collection.Cache
	FileModel   model.FileModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormDB, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表名前缀
			SingularTable: true, // 使用单数表
		},
		Logger: gormlogger.New(gormlogger.Config{LogLevel: logger.Info}),
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

	cache, err := collection.NewCache(time.Minute, collection.WithLimit(10000))
	Redis := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:      c,
		RedisClient: Redis,
		Cache:       cache,
		FileModel:   model.NewFileModel(gormDB),
	}
}
