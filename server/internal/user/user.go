package user

import (
	"context"
	"errors"
	"fmt"
	userInfo "user-info"
	"user-info/server/internal/repository"
)

type service struct {
	userDetailRepo repository.UserDetail
}

func NewService(userDetailRepo repository.UserDetail) *service {
	return &service{
		userDetailRepo: userDetailRepo,
	}
}

func (s service) GetUser(_ context.Context, request *userInfo.GetUserRequest) (*userInfo.GetUserResponse, error) {
	fmt.Println(request.Email)
	if len(request.Email) == 0 {
		return nil, errors.New("email is required")
	}

	//find user by email
	entity, err := s.userDetailRepo.FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	var response userInfo.GetUserResponse
	if entity == nil {
		fmt.Println("email is not found")
		return &response, nil
	}

	return &userInfo.GetUserResponse{
		Detail: &userInfo.User{
			Id:        entity.Id,
			Firstname: entity.Firstname,
			Lastname:  entity.Lastname,
			Email:     entity.Email,
		},
	}, nil
}

func (s service) GetAllUser(_ context.Context, _ *userInfo.GetAllUserRequest) (*userInfo.GetAllUserResponse, error) {
	entities, err := s.userDetailRepo.FindAll()
	if err != nil {
		return nil, err
	}
	users := initUsers(entities)
	return &userInfo.GetAllUserResponse{
		Details: users,
	}, nil
}

func initUsers(entities []repository.UserDetailEntity) []*userInfo.User {
	var tmp []*userInfo.User
	for _, val := range entities {
		tmp = append(tmp, &userInfo.User{
			Id:        val.Id,
			Firstname: val.Firstname,
			Lastname:  val.Lastname,
			Email:     val.Email,
		})
	}
	return tmp
}
