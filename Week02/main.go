package main

import (
	"fmt"
	"log"

	"github.com/ducknightii/Go-000/Week02/biz"
	"github.com/ducknightii/Go-000/Week02/dao"
)

func main() {
	dao.Init()

	age, err := biz.Search("haha")
	if err != nil {
		log.Fatalf("find age failed: %+v", err)
	}
	fmt.Println(age)
}
