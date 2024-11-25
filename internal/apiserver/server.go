package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/internal/model/server"
)

type apiServer struct {
	rest *server.REStfulAPIServer
}

type PreparedAPIServer struct {
	*apiServer
}

func New(cfg *model.Config) *apiServer {
	return &apiServer{
		rest: server.New(cfg),
	}
}

func (s *apiServer) Prepare() *PreparedAPIServer {
	r := gin.Default()
	InitRouter(r)

	// 创建 server 实例
	s.rest.Engine = r

	return &PreparedAPIServer{s}
}

func (s *PreparedAPIServer) Run() {
	s.apiServer.rest.Run()
}
