/*
 -- @author: shanweidi
 -- @date: 2023-04-18 11:14 上午
 -- @Desc:
*/
package sdk

type Protocol = string

const (
	MQTT  Protocol = "MQTT"
	MQTTS Protocol = "MQTTS"

	//设备注册
	PUBLISH_TOPIC_DEV_REG       = "hbt/iot/mbms/register"
	SUBSCRIBE_TOPIC_DEV_REG_RES = "hbt/iot/mbms/register/response"

	//网关上报运行数据
	PUBLISH_TOPIC_DEV_REPORT = "hbt/iot/mbms/report"

	//请求设备网关上报物模型数据
	SUBSCRIBE_TOPIC_DEV_GET        = "hbt/iot/mbms/get"
	PUBLISH_TOPIC_DEV_GET_DATA_RES = "hbt/iot/mbms/get/response"

	SUBSCRIBE_TOPIC_DEV_SET        = "hbt/iot/mbms/set"
	PUBLISH_TOPIC_DEV_SET_DATA_RES = "hbt/iot/mbms/set/response"

	PUBLISH_TOPIC_DEV_REQUEST = "hbt/iot/mbms/request"
	SUBSCRIBE_TOPIC_CLOUD_RES = "hbt/iot/mbms/request/response"
)
