/**
 * @Author: xueyanghan
 * @File: redis_subscribe.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/30 11:46
 */

package model

const (
	RedisSubscribeTableName = "t_redis_subscribe"
)

// RedisSubscribe -
type RedisSubscribe struct {
	CommonHeader

	StreamName    string `json:"StreamName" gorm:"column:stream_name;	type:varchar(128);	not null;	index;	comment:'流名称'"`
	GroupName     string `json:"GroupName" gorm:"column:group_name;	type:varchar(128);	not null;	index;	comment:'组名称'"`
	ConsumerName  string `json:"ConsumerName" gorm:"column:consumer_name;	type:varchar(128);	not null;	index;	comment:'消费者名称'"`
	MessageId     string `json:"MessageId" gorm:"column:message_id;	type:varchar(128);	not null;	index;	comment:'消息ID'"`
	MessageStatus int    `json:"MessageStatus" gorm:"column:message_status;	type:int8;	not null default 0;	index;	comment:'消息状态'; 0-未处理, 1-已处理"`

	CommonFooter
}

// TableName -
func (RedisSubscribe) TableName() string {
	return RedisSubscribeTableName
}
