/*
 -- @author: shanweidi
 -- @date: 2023-04-18 2:32 下午
 -- @Desc:
*/
package iBuilding

import "github.com/shanweidi/midea-BIoT-sdk-go/sdk"

type Client struct {
	sdk.Client
}

func NewClient() (client *Client, err error) {
	client = &Client{}
	err = client.InitClient()
	return
}

func NewClientWithOptions(config *sdk.Config) (client *Client, err error) {
	client = &Client{}
	err = client.InitClientWithConfig(config)
	return
}
