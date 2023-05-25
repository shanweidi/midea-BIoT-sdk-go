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
//seqNo 参数应等同于本次云端南向 Get 请求中所携带的seqNo
func (client *Client) ResponseGetRequest(op string, seqNo int, data interface{}) {
	if client.CouldAsync() {
		client.ResponseToEmqxAsync(sdk.PUBLISH_TOPIC_DEV_GET_DATA_RES, op, seqNo, data)
	} else {
		client.ResponseToEmqx(sdk.PUBLISH_TOPIC_DEV_GET_DATA_RES, op, seqNo, data)
	}
}

//ResponseSetRequest 用于回复 SubscribeSetRequest
//seqNo 参数应等同于本次云端南向 Set 命令中所携带的seqNo
func (client *Client) ResponseSetRequest(op string, seqNo int, data interface{}) {
	if client.CouldAsync() {
		client.ResponseToEmqxAsync(sdk.PUBLISH_TOPIC_DEV_SET_DATA_RES, op, seqNo, data)
	} else {
		client.ResponseToEmqx(sdk.PUBLISH_TOPIC_DEV_SET_DATA_RES, op, seqNo, data)
	}
}

//ReportRuntimeData 用于实时数据主动上报
//该 seqNo 由SDK用户自己维护，建议为随上报包数递增
func (client *Client) ReportRuntimeData(op string, seqNo int, data interface{}) {
	if client.CouldAsync() {
		client.PublishToEmqxAsync(sdk.PUBLISH_TOPIC_DEV_REPORT, op, seqNo, data)
	} else {
		client.PublishToEmqx(sdk.PUBLISH_TOPIC_DEV_REPORT, op, seqNo, data)
	}
}
