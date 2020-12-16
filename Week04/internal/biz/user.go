package biz

import "github.com/ducknightii/Go-000/Week04/internal/data"

type User struct {
	Name string `json:"name"`
	Age int32 `json:"age"`
}

func UserInfo(id int64) (user User, err error) {
	_user, err := data.UserInfo(id)
	user.Name = _user.Name
	user.Age = _user.Age
	return
}