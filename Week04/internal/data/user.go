package data

import (
	"github.com/ducknightii/Go-000/Week04/internal/pkg/db"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key;auto_increment"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name" gorm:"type:varchar(64);not null;default:'';INDEX:idx_name" comment:"name"`
	Age       int32  `json:"age" gorm:"type:smallint;not null;default:0" comment:"age"`
}

func UserInfo(id int64) (user User, err error) {
	err = db.DB.Where("id = ?", id).First(&user).Error
	return
}