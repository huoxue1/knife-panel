package app

import (
	"knife-panel/internal/app/config"
	"knife-panel/pkg/auth"
	"knife-panel/pkg/auth/jwtauth"
	"knife-panel/pkg/auth/jwtauth/store/buntdb"
	"knife-panel/pkg/auth/jwtauth/store/redis"
	jwt "github.com/dgrijalva/jwt-go"
)

// InitAuth 初始化用户认证
func InitAuth() (auth.Auther, error) {
	cfg := config.Global().JWTAuth

	var opts []jwtauth.Option
	opts = append(opts, jwtauth.SetExpired(cfg.Expired))
	opts = append(opts, jwtauth.SetSigningKey([]byte(cfg.SigningKey)))
	opts = append(opts, jwtauth.SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auth.ErrInvalidToken
		}
		return []byte(cfg.SigningKey), nil
	}))

	var method jwt.SigningMethod
	switch cfg.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}
	opts = append(opts, jwtauth.SetSigningMethod(method))

	var store jwtauth.Storer
	switch cfg.Store {
	case "redis":
		rcfg := config.Global().Redis
		store = redis.NewStore(&redis.Config{
			Addr:      rcfg.Addr,
			Password:  rcfg.Password,
			DB:        cfg.RedisDB,
			KeyPrefix: cfg.RedisPrefix,
		})
	default:
		s, err := buntdb.NewStore(cfg.FilePath)
		if err != nil {
			return nil, err
		}
		store = s
	}

	return jwtauth.New(store, opts...), nil
}
