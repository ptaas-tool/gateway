package jwt

import "time"

type Authenticator interface {
	GenerateToken(name string) (string, time.Time, error)
	ParseToken(token string) (string, error)
}

func New(cfg Config) Authenticator {
	return &authenticator{
		key:    cfg.PrivateKey,
		expire: cfg.ExpireTime,
	}
}
