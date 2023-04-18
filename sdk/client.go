/*
 -- @author: shanweidi
 -- @date: 2023-04-14 3:55 下午
 -- @Desc:
*/
package sdk

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/shanweidi/midea-BIoT-sdk-go/sdk/entities"
	"github.com/shanweidi/midea-BIoT-sdk-go/tools"
)

type Client struct {
	rw             sync.RWMutex
	config         *Config
	client         mqtt.Client
	LC             tools.LoggingClient
	isOpenAsync    bool
	asyncTaskQueue chan func()
}

func (client *Client) InitClientConfig() (config *Config) {
	if client.config != nil {
		return client.config
	} else {
		return NewConfig()
	}
}

func (client *Client) InitClient() tools.Error {
	client.rw.Lock()
	defer client.rw.Unlock()

	config := client.InitClientConfig()
	client.config = config
	return client.newMqttClient()
}

func (client *Client) Shutdown() {
	if client.asyncTaskQueue != nil {
		close(client.asyncTaskQueue)
	}

	client.isOpenAsync = false
}

// EnableAsync enable the async task queue
func (client *Client) EnableAsync(routinePoolSize, maxTaskQueueSize int) {
	if client.isOpenAsync {
		fmt.Println("warning: Please not call EnableAsync repeatedly")
		return
	}
	client.isOpenAsync = true
	client.asyncTaskQueue = make(chan func(), maxTaskQueueSize)
	for i := 0; i < routinePoolSize; i++ {
		go func() {
			for {
				task, notClosed := <-client.asyncTaskQueue
				if !notClosed {
					return
				} else {
					task()
				}
			}
		}()
	}
}

// registerResponseHandler 注册相应处理函数
func (client *Client) registerResponseHandler() mqtt.MessageHandler {
	return func(c mqtt.Client, message mqtt.Message) {
		// 报文解析
		mqttPayload := entities.CloudMqttBasicPayload{}
		err := json.Unmarshal(message.Payload(), &mqttPayload)
		if err != nil {
			client.LC.Errorf("register response handler json decode error, payload: %s", message.Payload())
			return
		}
		bytes, err := base64.StdEncoding.DecodeString(mqttPayload.Payload)
		if err != nil {
			client.LC.Errorf("invalid payload:%s", mqttPayload.Payload)
			return
		}
		mqttPayload.Payload = string(bytes)
		r := entities.BasicResponse{}
		if err = json.Unmarshal([]byte(mqttPayload.Payload), &r); err != nil {
			client.LC.Errorf("register response handler json decode error, payload: %s", mqttPayload.Payload)
			return
		}
		// 校验注册结果
		if r.Result != 0 {
			client.LC.Errorf("register response handler error. code: %s message: %s ", r.ErrCode, r.ErrMsg)
			return
		}

		//设备注册成功
		client.LC.Infof("----------register to cloud success----------")
		//todo 上报设备模型
		//go mqtt.onAuthPassedHandler(mqtt)
	}
}

// subscribeRegisterResponse 订阅注册相应
func (client *Client) subscribeRegisterResponse(c mqtt.Client, topic string) {
	token := c.Subscribe(topic, 0, client.registerResponseHandler())
	if token.Wait() && token.Error() != nil {
		client.LC.Errorf("subscribe register response error: %s", topic, token.Error())
	}
}

//建立起TCP连接后，网关先发送注册报文至云端，注册成功后再执行后续逻辑
func (client *Client) defaultOnConnectHandler() mqtt.OnConnectHandler {
	return func(c mqtt.Client) {
		client.subscribeRegisterResponse(c, tools.JoinMqttTopic(SUBSCRIBE_TOPIC_DEV_REG_RES, client.config.GwType, client.config.GwSn))

		//发送网关注册报文
		devRegPayload := entities.DevReg{
			Key:        client.config.Key,
			ReqTime:    strconv.Itoa(int(time.Now().UnixNano() / 1e6)),
			GwVer:      "0.0.1",
			ProductKey: "edgex",
			DevType:    "",
			DevSn:      client.config.GwSn,
			DevVer:     "devVer",
			MfrName:    "Midea",
			MfrModel:   "edgex",
		}
		//todo 这个seqNo待解决
		client.PublishToEmqx(PUBLISH_TOPIC_DEV_REG, "DEV_REG", 0, devRegPayload)
	}
}

func (client *Client) newMqttClient() tools.Error {
	opts := mqtt.NewClientOptions().AddBroker(client.config.ServerUri).SetClientID(client.config.ClientId)
	opts.SetKeepAlive(client.config.KeepAlive)
	opts.SetPassword(client.config.Password)
	opts.SetUsername(client.config.Username)
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(client.defaultOnConnectHandler())

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return tools.NewSdkError(tools.ConnectEmqxErrorCode, tools.ConnectEmqxErrorMessage, token.Error())
	}
	client.client = c
	return nil
}

func (client *Client) PublishToEmqx(topicPrefix, op string, seqNo int, data interface{}) {
	client.sendToBroker(tools.JoinMqttTopic(topicPrefix, client.config.GwType, client.config.GwSn, strconv.Itoa(seqNo)), op, seqNo, data)
}

func (client *Client) sendToBroker(topic string, op string, seqNo int, data interface{}) {
	p, _ := json.Marshal(data)
	client.LC.Infof("topic: %s ,op: %s ,seqNo: %d ,payload before base64: %s", topic, op, seqNo, p)
	mqttPayload := entities.CloudMqttBasicPayload{Op: op, SeqNo: seqNo, Payload: base64.StdEncoding.EncodeToString(p)}
	p, _ = json.Marshal(mqttPayload)

	token := client.client.Publish(topic, client.config.Qos, false, p)
	if token.Error() != nil {
		client.LC.Errorf("publish topic:%s message error:%v, payload:%s", topic, token.Error(), string(p))
		// 判断错误，是否重新初始化 MQTT clients
		if token.Error() == mqtt.ErrNotConnected {
			var err tools.Error
			err = client.newMqttClient()
			if err != nil {
				client.LC.Errorf("new mqtt client error: %s", err)
			}
		}
	}

}