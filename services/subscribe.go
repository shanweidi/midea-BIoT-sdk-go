/*
 -- @author: shanweidi
 -- @date: 2023-04-19 4:48 下午
 -- @Desc:
*/
package services

import (
	"github.com/shanweidi/midea-BIoT-sdk-go/sdk"
	"github.com/shanweidi/midea-BIoT-sdk-go/sdk/entities"
)

//SubscribeGetRequest 订阅来自云端的get请求
func (client *Client) SubscribeGetRequest(callback func(payload entities.CloudMqttBasicPayload)) {
	client.Subscribe(sdk.SUBSCRIBE_TOPIC_DEV_GET, callback)
}

//SubscribeSetRequest 订阅来自云端的set请求
func (client *Client) SubscribeSetRequest(callback func(payload entities.CloudMqttBasicPayload)) {
	client.Subscribe(sdk.SUBSCRIBE_TOPIC_DEV_SET, callback)
}
