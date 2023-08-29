/*
 -- @author: shanweidi
 -- @date: 2023-04-14 5:12 下午
 -- @Desc:
*/
package tools

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func InitStructWithDefaultTag(bean interface{}) {
	configType := reflect.TypeOf(bean)
	for i := 0; i < configType.Elem().NumField(); i++ {
		field := configType.Elem().Field(i)
		defaultValue := field.Tag.Get("default")
		if defaultValue == "" {
			continue
		}
		setter := reflect.ValueOf(bean).Elem().Field(i)
		switch field.Type.String() {
		case "int":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "uint8":
			intValue, _ := strconv.ParseUint(defaultValue, 10, 64)
			setter.SetUint(intValue)
		case "time.Duration":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue * 1000 * 1000 * 1000)
		case "string":
			setter.SetString(defaultValue)
		case "bool":
			boolValue, _ := strconv.ParseBool(defaultValue)
			setter.SetBool(boolValue)
		}
	}
}

func GetTimeInFormatISO8601() (timeStr string) {
	gmt := time.FixedZone("GMT", 0)

	return time.Now().In(gmt).Format("2006-01-02T15:04:05Z")
}

func JoinMqttTopic(elem ...string) string {
	return path.Join(elem...)
}

func ToJson(param interface{}) string {
	b, _ := json.Marshal(param)
	return string(b)
}

func ParseServerUri(uri string) string {
	if !strings.Contains(uri, "://") {
		return uri
	} else {
		return strings.Split(uri, "://")[1]
	}
}

func NewTlsConfig(filename string) *tls.Config {
	certpool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}
	certpool.AppendCertsFromPEM(ca)
	return &tls.Config{
		RootCAs: certpool,
	}
}
