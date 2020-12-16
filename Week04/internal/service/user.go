package service

import (
	"context"
	"errors"

	"github.com/ducknightii/Go-000/Week04/api"
	"github.com/ducknightii/Go-000/Week04/internal/biz"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/status"
)

type UserService struct {
}

func (s *UserService) UserInfo(ctx context.Context, req *api.UserRequest) (*api.UserResponse, error) {
	user, err := biz.UserInfo(req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// todo err enum pb定义
			return nil, status.Error(404, "user not found")
		}
	}
	return &api.UserResponse{
		Name: user.Name,
		Age:  user.Age,
	}, nil
}

func (s *UserService) mustEmbedUnimplementedUserServer() {

}
