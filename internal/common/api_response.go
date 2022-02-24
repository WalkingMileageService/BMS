package common

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func MakeHttpGetResponseToByte(component string, baseUrl string, param string) (responseByte []byte, err error) {
	now := RequestLog(component, echo.GET, baseUrl+param, "")

	request, err := http.NewRequest(echo.GET, baseUrl+param, nil)
	setHeader(request)
	responseByte, err = makeResponseByte(request)

	ResponseLog(component, echo.GET, baseUrl, string(responseByte), time.Since(now).Milliseconds())
	return
}

func MakeHttpPostResponseToByte(component string, baseUrl string, requestByte []byte) (responseByte []byte, err error) {
	now := RequestLog(component, echo.POST, baseUrl, string(requestByte))

	requestBody := bytes.NewBuffer(requestByte)
	request, err := http.NewRequest(echo.POST, baseUrl, requestBody)
	setHeader(request)

	responseByte, err = makeResponseByte(request)

	ResponseLog(component, echo.POST, baseUrl, string(responseByte), time.Since(now).Milliseconds())
	return
}

func MakeApiUrl(list []string, key string) string {
	queryParam := url.Values{}
	length := len(list)
	for i := 0; i < length; i++ {
		queryParam.Add(key, list[i])
	}
	return queryParam.Encode()
}

func makeResponseByte(request *http.Request) (responseByte []byte, err error) {
	client := &http.Client{}

	response, err := client.Do(request)
	if response != nil {
		responseByte, err = ioutil.ReadAll(response.Body)
	} else {
		return
	}

	responseByte, err = ioutil.ReadAll(response.Body)
	return
}
