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

	ContractInvalidErrorCode    = "ContractInvalid"
	ContractInvalidErrorMessage = "ContractInvalid, please check your param"
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
