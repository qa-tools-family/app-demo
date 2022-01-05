package grpc_server

import (
	"fmt"
	"github.com/qa-tools-family/app-demo/internal/options"
	"github.com/qa-tools-family/app-demo/internal/pb"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	*grpc.Server
	Address string
}


func (s *GrpcServer) Run() {
	listen, err := net.Listen("tcp", s.Address)
	if err != nil {
		logger.Fatalf("failed to listen: %s", err.Error())
	}

	go func() {
		if err := s.Serve(listen); err != nil {
			logger.Fatalf("failed to start grpc server: %s", err.Error())
		}
	}()

	logger.Infof("start grpc server at %s", s.Address)
}


func NewGrpcServer(o *options.GRPCOptions) (*GrpcServer, error) {
	// 根据 grpcOptions 生成对应的配置
	address := fmt.Sprintf("%s:%d", o.BindAddress, o.BindPort)
	opts := []grpc.ServerOption{grpc.MaxRecvMsgSize(4 * 1024 * 1024)}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterGreeterServer(grpcServer, &greeterServer{})

	return &GrpcServer{Server: grpcServer, Address: address}, nil
}
