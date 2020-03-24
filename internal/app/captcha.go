package app

import (
	"github.com/LyricTian/captcha"
	"github.com/LyricTian/captcha/store"
	"knife-panel/internal/app/config"
	"knife-panel/pkg/logger"
	"github.com/go-redis/redis"
)

// InitCaptcha 初始化图形验证码
func InitCaptcha() {
	cfg := config.Global().Captcha
	if cfg.Store == "redis" {
		rc := config.Global().Redis
		captcha.SetCustomStore(store.NewRedisStore(&redis.Options{
			Addr:     rc.Addr,
			Password: rc.Password,
			DB:       cfg.RedisDB,
		}, captcha.Expiration, logger.StandardLogger(), cfg.RedisPrefix))
	}
}
