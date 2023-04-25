/*
 -- @author: shanweidi
 -- @date: 2023-04-14 3:53 下午
 -- @Desc:
*/
package sdk

import (
	"time"

	"github.com/shanweidi/midea-BIoT-sdk-go/tools"
)

type Config struct {
	GwType            string `default:""`
	GwSn              string `default:""`
	Key               string `default:""`
	ProductKey        string `default:""`
	ServerUri         string `default:""`
	ClientId          string `default:""`
	Protocol          string `default:"MQTT"`
	GoRoutinePoolSize int    `default:"5"`
	MaxTaskQueueSize  int    `default:"1000"`
	EnableAsync       bool   `default:"false"`
	Qos               byte   `default:"0"`
	Username          string `default:""`
	Password          string `default:""`
	LogLevel          string `default:"INFO"`
	Timeout           time.Duration
	KeepAlive         time.Duration `default:"60"`
}

func NewConfig() (config *Config) {
	config = &Config{}
	tools.InitStructWithDefaultTag(config)
	return
}

func (c *Config) WithEnableAsync(isEnableAsync bool) *Config {
	c.EnableAsync = isEnableAsync
	return c
}

func (c *Config) WithMaxTaskQueueSize(maxTaskQueueSize int) *Config {
	c.MaxTaskQueueSize = maxTaskQueueSize
	return c
}

func (c *Config) WithGoRoutinePoolSize(goRoutinePoolSize int) *Config {
	c.GoRoutinePoolSize = goRoutinePoolSize
	return c
}

func (c *Config) WithProtocol(protocol string) *Config {
	c.Protocol = protocol
	return c
}

func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}

func (c *Config) WithKeepAlive(keepalive time.Duration) *Config {
	c.KeepAlive = keepalive
	return c
}

func (c *Config) WithLogLevel(level string) *Config {
	c.LogLevel = level
	return c
}

func (c *Config) WithGwSn(gwSn string) *Config {
	c.GwSn = gwSn
	return c
}

func (c *Config) WithGwType(gwType string) *Config {
	c.GwType = gwType
	return c
}

func (c *Config) WithKey(key string) *Config {
	c.Key = key
	return c
}

func (c *Config) WithProductKey(productKey string) *Config {
	c.ProductKey = productKey
	return c
}

func (c *Config) WithServerUri(uri string) *Config {
	c.ServerUri = uri
	return c
}

func (c *Config) WithClientId(id string) *Config {
	c.ClientId = id
	return c
}

func (c *Config) WithQos(qos byte) *Config {
	c.Qos = qos
	return c
}

func (c *Config) WithUsername(username string) *Config {
	c.Username = username
	return c
}

func (c *Config) WithPassword(password string) *Config {
	c.Password = password
	return c
}
