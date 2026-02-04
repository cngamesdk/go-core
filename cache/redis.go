package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

type RedisConfig struct {
	Db       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Prefix   string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
}

type Client struct {
	client *redis.Client
	config RedisConfig
	logger *zap.Logger
}

func NewClient(config RedisConfig, logger *zap.Logger) (*Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
		PoolSize: 100,
	})

	// 测试连接
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Client{
		client: rdb,
		config: config,
		logger: logger,
	}, nil
}

// 缓存键定义
func (c *Client) key(parts ...string) string {
	key := c.config.Prefix
	for _, part := range parts {
		key += ":" + part
	}
	return key
}

// 限流和计数器
func (c *Client) RateLimit(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	current, err := c.client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if current == 1 {
		_ = c.client.Expire(ctx, key, window).Err()
	}

	return current > int64(limit), nil
}

// 分布式锁
func (c *Client) AcquireLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return c.client.SetNX(ctx, key, "1", ttl).Result()
}

func (c *Client) ReleaseLock(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// OpenRedis 打开Redis链接
func OpenRedis(config RedisConfig) (resp *redis.Client, err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.Db,       // use default DB
	})
	resp = rdb
	return
}
