/**
 * @Author: xueyanghan
 * @File: redis_subscribe_types.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/5/11 11:46
 */

package types

// UserInfo -
type UserInfo struct {
	UserName string `json:"UserName" comment:"用户名"`
	Password string `json:"Password" comment:"密码, md5加密"`
	Email    string `json:"Email" comment:"邮箱"`
	Phone    string `json:"Phone" comment:"手机号"`
	Role     int    `json:"Role" comment:"角色, 0-普通用户, 1-管理员"`
}
