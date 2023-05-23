/*
 -- @author: shanweidi
 -- @date: 2023-05-06 2:43 下午
 -- @Desc:
*/
package sdk

import (
	"fmt"
	"testing"

	"github.com/shanweidi/midea-BIoT-sdk-go/sdk/entities"
	"github.com/stretchr/testify/assert"
)

func Test_InitClientWithConfig(t *testing.T) {
	config := NewConfig().WithClientId("shan_test").WithServerUri("mqtt://127.0.0.1:1883")
	client := &Client{}
	err := client.InitClientWithConfig(config)
	assert.Nil(t, err)
	assert.False(t, client.isOpenAsync)
}

func Test_EnableAsync(t *testing.T) {
	config := NewConfig().WithGoRoutinePoolSize(1).WithMaxTaskQueueSize(2)
	client := &Client{}
	err := client.InitClientWithConfig(config)
	asyncErr := client.AddAsyncTask(func() {})
	assert.Nil(t, asyncErr)
	client.EnableAsync(10, 10)
	assert.True(t, client.isOpenAsync)
	assert.Nil(t, err)
	assert.NotNil(t, client)

	client.Shutdown()
	assert.False(t, client.isOpenAsync)
}

func Test_Subscribe(t *testing.T) {
	config := NewConfig().WithClientId("shan_test").WithServerUri("mqtt://127.0.0.1:1883")
	client := &Client{}
	client.InitClientWithConfig(config)

	client.Subscribe("test", func(basicPayload entities.CloudMqttBasicPayload) {
		fmt.Println("subscribe test")
	})
	select {}
}
