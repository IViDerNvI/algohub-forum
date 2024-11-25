package server

import "time"

type JwtInfo struct {
	Realm      string
	Key        string
	Timeout    time.Duration
	MaxRefresh time.Duration
}
