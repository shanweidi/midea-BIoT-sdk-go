/*
 -- @author: shanweidi
 -- @date: 2023-04-18 10:13 上午
 -- @Desc:
*/
package tools

import "fmt"

const (
	DefaultSdkErrorCode = "SDK.Error"

	ConnectEmqxErrorCode    = "SDK.ConnectEmqxFail"
	ConnectEmqxErrorMessage = "Connect Emqx Broker failed, please check your connect options"

	RegisterKeyAuthErrorCode    = "SDK.KeyAuthFail"
	RegisterKeyAuthErrorMessage = "RegisterKey is not correct, please check register payload"

	ContractInvalidErrorCode    = "SDK.ContractInvalid"
	ContractInvalidErrorMessage = "ContractInvalid, please check your param"

	AsyncFunctionNotEnabledCode    = "SDK.AsyncFunctionNotEnabled"
	AsyncFunctionNotEnabledMessage = "Async function is not enabled in client, please invoke 'client.EnableAsync' function"

	ConfigServerUrlErrorCode    = "SDK.ServerUriEmpty"
	ConfigServerUrlErrorMessage = "Your Emqx Server Uri is empty, please check your config"

	ConfigClientIdErrorCode    = "SDK.ClientIdEmpty"
	ConfigClientIdErrorMessage = "Your Client ID is empty, please check your config"

	ConfigGatewaySnErrorCode    = "SDK.GatewaySnEmpty"
	ConfigGatewaySnErrorMessage = "Your Gateway SN is empty, please check your config"

	ConfigGatewayTypeErrorCode    = "SDK.GatewayTypeEmpty"
	ConfigGatewayTypeErrorMessage = "Your Gateway Type is empty, please check your config"

	ConfigKeyErrorCode    = "SDK.SecretKeyEmpty"
	ConfigKeyErrorMessage = "Your Secret Key is empty, please contact midea BIoT acquire!"
)

type Error interface {
	error
	ErrorCode() string
	Message() string
	OriginError() error
}

type SdkError struct {
	originError error
	errorCode   string
	message     string
}

func NewSdkError(errorCode, message string, originErr error) Error {
	return &SdkError{
		errorCode:   errorCode,
		message:     message,
		originError: originErr,
	}
}

func (err *SdkError) Error() string {
	clientErrMsg := fmt.Sprintf("[%s] %s", err.ErrorCode(), err.message)
	if err.originError != nil {
		return clientErrMsg + "\ncaused by:\n" + err.originError.Error()
	}
	return clientErrMsg
}

func (err *SdkError) ErrorCode() string {
	if err.errorCode == "" {
		return DefaultSdkErrorCode
	} else {
		return err.errorCode
	}
}

func (err *SdkError) Message() string {
	return err.message
}

func (err *SdkError) OriginError() error {
	return err.originError
}

func (err *SdkError) String() string {
	return err.Error()
}
