package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	userInfo "user-info"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer cc.Close()
	service := userInfo.NewUserInfoServiceClient(cc)

	//get user by email
	fmt.Println("start get user by email")
	getUserByEmail(service)
	fmt.Printf("\n=====================================\n")

	//get all user by id
	fmt.Println("start get all users")
	getAllUser(service)
}

func getUserByEmail(service userInfo.UserInfoServiceClient) {
	resp, err := service.GetUser(context.Background(), &userInfo.GetUserRequest{
		Email: "rkritchat@gmail.com",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%#v", resp.Detail)
}

func getAllUser(service userInfo.UserInfoServiceClient) {
	resp, err := service.GetAllUser(context.Background(), &userInfo.GetAllUserRequest{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("total user is %v\n", len(resp.Details))
	for _, val := range resp.Details {
		fmt.Printf("%#v", val)
	}
}
