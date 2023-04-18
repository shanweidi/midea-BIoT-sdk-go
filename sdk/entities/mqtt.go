/*
 -- @author: shanweidi
 -- @date: 2023-04-18 3:04 下午
 -- @Desc:
*/
package entities

type CloudMqttBasicPayload struct {
	Op       string `json:"op"`
	SeqNo    int    `json:"seqNo"`
	Encoding string `json:"encoding"`
	Payload  string `json:"payload"`
}

type BasicResponse struct {
	Result  int    `json:"result"`
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

// 设备注册报文
type DevReg struct {
	Key        string `json:"key"`
	ReqTime    string `json:"reqTime"`
	GwVer      string `json:"gwVer"`
	ProductKey string `json:"productKey"`
	DevType    string `json:"devType"`
	DevSn      string `json:"devSn"`
	DevVer     string `json:"devVer"`
	MfrName    string `json:"mfrName"`
	MfrModel   string `json:"mfrModel"`
}
