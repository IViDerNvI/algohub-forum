package server

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub-forum/internal/apiserver/middleware"
	"github.com/ividernvi/algohub-forum/internal/apiserver/pkg/core"
	"github.com/ividernvi/algohub-forum/internal/model"
	"github.com/ividernvi/algohub-forum/pkg/code"
	"github.com/sirupsen/logrus"
)

type REStfulAPIServer struct {
	middlewares []string

	SecureServingInfo   *SecureServingInfo
	InsecureServingInfo *InsecureServingInfo

	*gin.Engine
	healthz bool

	insecureServer, secureServer *http.Server
}

func New(cfg *model.Config) *REStfulAPIServer {
	return &REStfulAPIServer{
		middlewares: []string{},

		SecureServingInfo: &SecureServingInfo{
			BindAddress: cfg.SecureServing.BindAddress,
			BindPort:    cfg.SecureServing.BindPort,
			ServerCert:  CertKey(cfg.SecureServing.ServerCert),
		},

		InsecureServingInfo: &InsecureServingInfo{
			BindAddress: cfg.InsecureServing.BindAddress,
			BindPort:    cfg.InsecureServing.BindPort,
		},

		Engine: gin.Default(),

		healthz: true,
	}
}

func (s *REStfulAPIServer) initRESTfulAPIServer() {
	s.setup()
	s.installMiddleware()
	s.installAPI()
}

func (s *REStfulAPIServer) setup() {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		logrus.Printf("Method: %s, Path: %s, Handler: %s (nuHandlers: %d)", httpMethod, absolutePath, handlerName, nuHandlers, nuHandlers)
	}
}

func (s *REStfulAPIServer) installMiddleware() {
	for _, m := range s.middlewares {
		middleware.Install(m, s.Engine)
	}
}

func (s *REStfulAPIServer) installAPI() {
	if s.healthz {
		s.GET("/healthz", func(ctx *gin.Context) {
			core.WriteResponse(ctx, code.ERROR_OK, map[string]string{"status": "ok"})
		})
	}

	s.GET("/version", func(ctx *gin.Context) {
		core.WriteResponse(ctx, code.ERROR_OK, model.GetVersion())
	})
}

func (s *REStfulAPIServer) Run() {
	s.insecureServer = &http.Server{
		Addr:    ":8089",
		Handler: s,
	}

	s.secureServer = &http.Server{
		Addr:    ":8090",
		Handler: s,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		if err := s.insecureServer.ListenAndServe(); err != nil {
			logrus.Fatalf("insecure server not run: %s", err)
		}
	}()

	wg.Add(1)
	go func() {
		crt := s.SecureServingInfo.ServerCert.CertFile
		key := s.SecureServingInfo.ServerCert.KeyFile

		logrus.Printf("%s, ", key)

		if err := s.secureServer.ListenAndServeTLS(crt, key); err != nil {
			logrus.Fatalf("secure server not run: %s", err)
		}
	}()

	wg.Wait()

	return
}
