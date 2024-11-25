package options

import "github.com/gin-gonic/gin"

type ServerRunOptions struct {
	Mode        string   `json:"mode"        mapstructure:"mode"`
	Healthz     bool     `json:"healthz"     mapstructure:"healthz"`
	Middlewares []string `json:"middlewares" mapstructure:"middlewares"`
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		Mode:        gin.ReleaseMode,
		Healthz:     true,
		Middlewares: []string{},
	}
}
