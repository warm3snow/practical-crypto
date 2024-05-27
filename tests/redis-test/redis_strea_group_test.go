/**
 * @Author: xueyanghan
 * @File: redis_strea_group_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/30 10:26
 */

package redis_test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strings"
	"testing"
)

func testStreamGroupXRead(done chan bool) {
	redisCli := newRedisCli()
	defer redisCli.Close()

	// 创建消费者组
	_, err := redisCli.XGroupCreateMkStream(context.Background(), streamName, groupName, "0").Result()
	if err != nil && !strings.Contains(err.Error(), "BUSYGROUP Consumer Group name already exists") {
		log.Printf("XGroupCreateMkStream failed: %v", err)
		return
	}
	// 添加消费者到组中
	_, err = redisCli.XGroupSetID(context.Background(), streamName, groupName, "1714447212112-0").Result()
	if err != nil {
		log.Fatal("Error adding consumer to group:", err)
	}


	// 读取消息
	//for {
		streams, err := redisCli.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Group:    groupName,
			Consumer: consumerName,
			Streams:  []string{streamName, ">"},
			//Count:    1,
			Block:    0,
		}).Result()

		if err != nil {
			log.Printf("XReadGroup failed: %v", err)
			return
		}


		for _, stream := range streams {
			if stream.Stream == streamName {
				for _, message := range stream.Messages {
					log.Printf("XReadGroup success, messageId: %s, values: %#v", message.ID,message.Values)
					redisCli.XAck(context.Background(), streamName, groupName, message.ID)
				}
			}
		}
	//}

	done <- true
}

func TestStreamGroup(t *testing.T) {
	done := make(chan bool, 1)

	go func() {
		testStreamGroupXRead(done)
	}()

	//go func() {
	//	testStreamXAdd(done, t)
	//}()

	<-done
}