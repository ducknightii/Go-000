package main

import (
	"github.com/ducknightii/Go-000/Week04/api"
	"github.com/ducknightii/Go-000/Week04/configs"
	"github.com/ducknightii/Go-000/Week04/internal/pkg/db"
	"github.com/ducknightii/Go-000/Week04/internal/service"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	// 初始化配置信息
	configs.InitConfig()
	// 初始化数据库
	db.Init()

	listen, err := net.Listen("tcp", configs.Conf.Server.Listen)
	if err != nil {
		os.Exit(1)
	}

	userService := service.UserService{}
	server := grpc.NewServer()
	api.RegisterUserServer(server, userService)
	if err := server.Serve(listen); err != nil {
		log.Fatalf("RPC server listen failed. err: %s\n", err.Error())
	}
}

