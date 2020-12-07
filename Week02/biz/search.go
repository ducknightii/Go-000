package biz

import "github.com/ducknightii/Go-000/Week02/dao"

func Search(name string) (int, error) {
	return dao.Age(name)
}
