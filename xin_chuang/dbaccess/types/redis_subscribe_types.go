/**
 * @Author: xueyanghan
 * @File: redis_subscribe_types.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/5/11 11:46
 */

package types

// RedisSubscribe -
type RedisSubscribe struct {
	StreamName    string `json:"StreamName" comment:"流名称"`
	GroupName     string `json:"GroupName" comment:"组名称"`
	ConsumerName  string `json:"ConsumerName" comment:"消费者名称"`
	MessageId     string `json:"MessageId" comment:"消息ID"`
	MessageStatus int    `json:"MessageStatus" comment:"消息状态, 0-未处理, 1-已处理"`
}
