/*
 -- @author: shanweidi
 -- @date: 2023-04-18 11:14 上午
 -- @Desc:
*/
package sdk

const (
	GW_TYPE = "edgex"
	//设备注册
	PUBLISH_TOPIC_DEV_REG       = "hbt/iot/mbms/register"
	SUBSCRIBE_TOPIC_DEV_REG_RES = "hbt/iot/mbms/register/response"
	//网关上报运行数据
	PUBLISH_TOPIC_DEV_REPORT = "hbt/iot/mbms/report/" + GW_TYPE
	//请求设备网关上报物模型数据
	SUBSCRIBE_TOPIC_DEV_GET        = "hbt/iot/mbms/get/" + GW_TYPE
	PUBLISH_TOPIC_DEV_GET_DATA_RES = "hbt/iot/mbms/get/response/" + GW_TYPE
	SUBSCRIBE_TOPIC_DEV_SET        = "hbt/iot/mbms/set/" + GW_TYPE
	PUBLISH_TOPIC_DEV_SET_DATA_RES = "hbt/iot/mbms/set/response/" + GW_TYPE
)
