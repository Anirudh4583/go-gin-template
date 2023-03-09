package utils

import (
	"time"
)

type Maker interface {
    CreateToken(username string, password string ,duration time.Duration) (string, error)
    VerifyToken(token string) (*Payload, error)
}