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

//SubscribeGetRequest 订阅来自云端的 Get 请求
//SDK用户自定义实现 callback
func (client *Client) SubscribeGetRequest(callback func(payload entities.CloudMqttBasicPayload)) {
	client.Subscribe(sdk.SUBSCRIBE_TOPIC_DEV_GET, callback)
}

//SubscribeSetCommand 订阅来自云端的 Set 命令
//SDK用户自定义实现 callback
func (client *Client) SubscribeSetCommand(callback func(payload entities.CloudMqttBasicPayload)) {
	client.Subscribe(sdk.SUBSCRIBE_TOPIC_DEV_SET, callback)
}

//SubscribeCloudResponse 用于订阅 ReportRequestData 后，来自云端的响应
//payload中的 seqNo 与 SDK用户在 ReportRequestData 中上报时的包序号保持一致
//SDK用户自定义实现 callback
func (client *Client) SubscribeCloudResponse(callback func(payload entities.CloudMqttBasicPayload)) {
	client.Subscribe(sdk.SUBSCRIBE_TOPIC_CLOUD_RES, callback)
}
