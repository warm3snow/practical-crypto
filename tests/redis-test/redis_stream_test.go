/**
 * @Author: xueyanghan
 * @File: redis_stream_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2024/4/30 10:07
 */

package redis_test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
	"time"
)

var(
	streamName = "mystream"
	groupName = "mygroup"
	consumerName = "myconsumer"
)

func newRedisCli() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "123456",         // Redis密码，如果没有设置密码则为空
		DB:       0,                // Redis数据库索引
	})
	return client
}

func testStreamXAdd(done chan bool, t *testing.T) {
	redisCli := newRedisCli()
	defer redisCli.Close()

	// create a stream
	for i := 0; i < 10; i++ {
		_, err := redisCli.XAdd(context.Background(), &redis.XAddArgs{
			Stream: streamName,
			Values: map[string]interface{}{
				fmt.Sprintf("key%d", i): fmt.Sprintf("value%d", i),
			},
		}).Result()

		if err != nil {
			log.Printf("XAdd failed: %v", err)
			return
		}

		log.Printf("XAdd success, key: %s, value: %s", fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))

		time.Sleep(1 * time.Second)
	}

	done <- true
}

func testStreamXRead(done chan bool,t *testing.T) {
	redisCli := newRedisCli()
	defer redisCli.Close()

	streams, err := redisCli.XRead(context.Background(), &redis.XReadArgs{
		Streams: []string{streamName, "0"},
		Count:   10,
		Block:   0,
	}).Result()

	if err != nil {
		log.Printf("XRead failed: %v", err)
		return
	}

	for _, stream := range streams {
		if stream.Stream == streamName {
			for _, message := range stream.Messages {
				log.Printf("XRead success, values: %#v", message.Values)
				redisCli.XAck(context.Background(), streamName, "", message.ID)
			}
		}
	}


	done <- true
}

func TestStreamXAddRead(t *testing.T) {
	done := make(chan bool, 2)

	go func() {
		testStreamXAdd(done, t)
	}()

	//go func() {
	//	testStreamXRead(done, t)
	//}()

	for i := 0; i < 2; i++ {
		<-done
	}
}