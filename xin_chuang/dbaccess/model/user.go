/**
 * @Author: xueyanghan
 * @File: redis_subscribe.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/30 11:46
 */

package model

const (
	UserTableName = "t_user"
)

// User -
type User struct {
	CommonHeader

	UserName string `gorm:"column:user_name; type:varchar(256); comment:'用户名'"`
	Password string `gorm:"column:password; type:varchar(256); comment:'密码, md5加密'"`
	Email    string `gorm:"column:email; type:varchar(256); comment:'邮箱'"`
	Phone    string `gorm:"column:phone; type:varchar(256); comment:'手机号'"`
	Role     int    `gorm:"column:role; type: int; comment:'角色, 0-普通用户, 1-管理员'"`

	CommonFooter
}

// TableName -
func (User) TableName() string {
	return UserTableName
}
