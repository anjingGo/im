package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint64 `gorm:"primaryKey" json:"id" form:"id"`
	UserName  string `json:"user_name" gorm:"comment:用户名"`
	Password  string `json:"-" gorm:"comment:用户密码"`
	Avatar    string `json:"avatar" gorm:"comment:头像"`
	NickName  string `json:"nick_name" gorm:"default:匿名;comment:用户昵称"`
	Phone     string `json:"phone" gorm:"comment:手机号"`
	Online    uint   `json:"online" gorm:"comment:在线状态"`
	Token     string `json:"token" gorm:"comment:token"`
	Role      string `json:"role" gorm:"comment:角色"`
	Status    uint   `json:"status" gorm:"comment:状态"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at"`
}
