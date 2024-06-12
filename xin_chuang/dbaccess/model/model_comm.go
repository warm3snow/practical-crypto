/**
 * @Author: xueyanghan
 * @File: model_comm.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/30 11:48
 */

package model

import "time"

type CommonHeader struct {
	ID int `json:"-" gorm:"column:id; primary_key; auto_increment; comment:'主键'"`
}

type CommonFooter struct {
	CreatedAt time.Time  `json:"-" sql:"index"    gorm:"column:create_time;    comment:'创建时间'"`
	UpdatedAt time.Time  `json:"-" sql:"index"    gorm:"column:modify_time;    comment:'修改时间'"`
	DeletedAt *time.Time `json:"-" sql:"index"    gorm:"column:delete_time;    comment:'删除时间'"`
}
