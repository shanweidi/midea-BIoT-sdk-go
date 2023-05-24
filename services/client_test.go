/*
 -- @author: shanweidi
 -- @date: 2023-05-23 2:44 下午
 -- @Desc:
*/
package services

import (
	"testing"

	"github.com/shanweidi/midea-BIoT-sdk-go/sdk"
	"github.com/stretchr/testify/assert"
)

func Test_Client(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
		assert.Equal(t, "not support yet", err)
	}()
	NewClient()
}

func Test_ClientWithConfig(t *testing.T) {
	config := sdk.NewConfig().WithServerUri("mqtt://127.0.0.1:1883")
	_, err := NewClientWithOptions(config)
	assert.NotNil(t, err)
	config.WithClientId("iBuilding").WithGwType("edgex").
		WithGwSn("6666").WithKey("uiowem9682")

	_, err = NewClientWithOptions(config)
	assert.Nil(t, err)
	select {}
}
