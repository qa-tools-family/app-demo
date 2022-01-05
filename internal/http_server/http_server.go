package http_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qa-tools-family/app-demo/internal/options"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

func defaultMiddlewares() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"recovery":  gin.Recovery(),
		"logger":    gin.Logger(),
	}
}

type HttpServer struct {
	*gin.Engine
	*http.Server
	middlewares []string
	mode        string
	address     string
}

func (s *HttpServer) Setup()  {
	s.Engine = gin.New()
	gin.SetMode(s.mode)
}

func (s *HttpServer) InstallMiddlewares()  {
	middlewares := defaultMiddlewares()
	for _, m := range s.middlewares {
		mw, ok := middlewares[m]
		if !ok {
			logger.Warnf("can not find middleware: %s", m)
			continue
		}
		logger.Infof("install middleware: %s", m)
		s.Use(mw)
	}
}

func (s *HttpServer) Run() {
	s.Server = &http.Server{
		Addr: s.address,
		Handler: s,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logger.Fatalf("failed to start http server: %s", err.Error())
		}
	}()

	logger.Infof("start http server at %s", s.address)
}

func initServer(g *HttpServer) {
	g.Setup()
	g.InstallMiddlewares()
	installController(g.Engine)
}

func NewHttpServer(o *options.HttpOptions) (*HttpServer, error) {
	// 根据 HttpOptions 生成对应的配置
	address := fmt.Sprintf("%s:%d", o.BindAddress, o.BindPort)
	httpServer := &HttpServer{address: address, mode: o.Mode, middlewares: o.Middlewares}
	initServer(httpServer)
	return httpServer, nil
}
