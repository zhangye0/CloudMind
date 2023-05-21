package svc

import (
	"CloudMind/app/es/cmd/rpc/internal/config"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	Es     *elasticsearch.Client
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	es, _ := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	})

	return &ServiceContext{
		Config: c,
		Es:     es,
		Redis:  redis.MustNewRedis(c.RedisConf),
	}
}
