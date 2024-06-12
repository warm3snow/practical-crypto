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

type ChainInfo struct {
	// ChainInfoId is from chain_service, don't generate automatically
	ChainInfoId int    `json:"-" gorm:"column:chain_info_id; type: int; comment:'链信息ID'"`
	ChainName   string `json:"-" gorm:"column:chain_name; type:varchar(256); comment:'链名称'"`
}

type TxInfo struct {
	TxId string `json:"TxId" gorm:"column:tx_id; type:varchar(128); comment:'交易ID'"`
}
