package config

import (
	"time"
)

var (
	JWTSecretKey           = []byte("emreskey")
	AccessTokenExpiration  = time.Minute * 15
	RefreshTokenExpiration = time.Hour * 7
)
