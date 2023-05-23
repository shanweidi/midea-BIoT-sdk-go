/*
 -- @author: shanweidi
 -- @date: 2023-04-19 4:48 下午
 -- @Desc:
*/
package services

import (
	"github.com/shanweidi/midea-BIoT-sdk-go/sdk"
)

//ResponseGetRequest 用于回复 SubscribeGetRequest
func (client *Client) ResponseGetRequest(op string, seqNo int, data interface{}) {
	if client.CouldAsync() {
		client.PublishToEmqxAsync(sdk.PUBLISH_TOPIC_DEV_GET_DATA_RES, op, seqNo, data)
	} else {
		client.PublishToEmqx(sdk.PUBLISH_TOPIC_DEV_GET_DATA_RES, op, seqNo, data)
	}
}

//ResponseSetRequest 用于回复 SubscribeSetRequest
func (client *Client) ResponseSetRequest(op string, seqNo int, data interface{}) {
	if client.CouldAsync() {
		client.PublishToEmqxAsync(sdk.PUBLISH_TOPIC_DEV_SET_DATA_RES, op, seqNo, data)
	} else {
		client.PublishToEmqx(sdk.PUBLISH_TOPIC_DEV_SET_DATA_RES, op, seqNo, data)
	}
}

func (client *Client) ReportRuntimeData(op string, seqNo int, data interface{}) {
	if client.CouldAsync() {
		client.PublishToEmqxAsync(sdk.PUBLISH_TOPIC_DEV_REPORT, op, seqNo, data)
	} else {
		client.PublishToEmqx(sdk.PUBLISH_TOPIC_DEV_REPORT, op, seqNo, data)
	}
}
