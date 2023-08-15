package jwt

import (
	"time"

	"github.com/apt-tool/apt-core/pkg/enum"
)

type Authenticator interface {
	GenerateToken(name string, role enum.Role) (string, time.Time, error)
	ParseToken(token string) (string, error)
}

func New(cfg Config) Authenticator {
	return &authenticator{
		key:    cfg.PrivateKey,
		expire: cfg.ExpireTime,
	}
}
