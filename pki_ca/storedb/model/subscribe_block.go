/**
 * @Author: xueyanghan
 * @File: subscribe_block.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/8/18 15:14
 */

package model

import "time"

var (
	SubscribeBlockTableName = "t_subscribe_block"

	ChainMakerChainName = "chainmaker"
)

// SubscribeBlock -
type SubscribeBlock struct {
	ID int `json:"-" gorm:"column:Fid;	type:int; 	primaryKey;	autoIncrement; 	comment:'自增ID'"`

	ChainName string `json:"ChainName" gorm:"column:Fchain_name;	type:varchar(128);	not null;	index;	comment:'链名称'"`
	Height    int64  `json:"Height" gorm:"column:Fheight;	type:bigint(20);	not null;	index;	comment:'区块高度'"`

	CreatedAt time.Time  `json:"-" sql:"index"    gorm:"column:Fcreate_time;    comment:'创建时间'"`
	UpdatedAt time.Time  `json:"-" sql:"index"    gorm:"column:Fmodify_time;    comment:'修改时间'"`
	DeletedAt *time.Time `json:"-" sql:"index"    gorm:"column:Fdelete_time;    comment:'删除时间'"`
}

// TableName -
func (SubscribeBlock) TableName() string {
	return SubscribeBlockTableName
}
