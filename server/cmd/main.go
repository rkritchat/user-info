package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	userInfo "user-info"
	"user-info/server/internal/config"
	"user-info/server/internal/repository"
	"user-info/server/internal/user"
)

func main() {
	//init config
	cfg := config.InitConfig()
	defer cfg.Free()

	//init repository
	userDetailRepo := repository.NewUserDetail(cfg.DB)
	//init service
	service := user.NewService(userDetailRepo)

	fmt.Println("start on port 50051")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln(err)
		return
	}

	s := grpc.NewServer()
	userInfo.RegisterUserInfoServiceServer(s, service)
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
