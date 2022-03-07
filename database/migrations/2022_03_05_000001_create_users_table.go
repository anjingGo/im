package migrations

import (
	"github.com/anjingGo/im/app/models"
)

func init() {
	type User struct {
		models.BaseModel

		UserName  string `gorm:"type:varchar(255);comment:用户名"`
		Password  string `gorm:"type:varchar(255);comment:用户密码"`
		Avatar    string `gorm:"comment:头像"`
		NickName  string `gorm:"default:匿名;comment:用户昵称"`
		Phone     string `gorm:"type:varchar(20);comment:手机号"`
		Online    uint   `gorm:"comment:在线状态"`
		Token     string `gorm:"comment:token"`
		Role      string `gorm:"comment:角色"`
		Status    uint   `gorm:"comment:状态"`

		models.CommonTimestampsField

		models.DeletedAtTimestampsField

	}
}