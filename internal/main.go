package main

import (
	"github.com/qa-tools-family/app-demo/internal/grpc_server"
	"github.com/qa-tools-family/app-demo/internal/http_server"
	"github.com/qa-tools-family/app-demo/internal/options"
	"github.com/qa-tools-family/go-common-kit/app"
	logger "github.com/sirupsen/logrus"
	"time"
)

func Run(opts *options.Options) error {
	grpcServer, err := grpc_server.NewGrpcServer(opts.GRPCOptions)
	if err != nil {
		logger.Errorf("new grpc server error: %s", err)
		return err
	}
	grpcServer.Run()

	httpServer, err := http_server.NewHttpServer(opts.HttpOptions)
	if err != nil {
		logger.Errorf("new http server error: %s", err)
		return err
	}
	httpServer.Run()
	time.Sleep(365 * 24 * time.Hour)
	return nil  // 启动
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		return Run(opts)
	}
}

const commandDesc = `This is App Demo Project`

// NewApp creates a App object with default parameters.
func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("App Demo",
		basename,
		app.WithOptions(opts),  // app 包要求 opts 自定义结构体需要实现 Flags 和 Validate 两个方法
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func main() {
	basename := "app-demo"
	NewApp(basename).Run()
}
