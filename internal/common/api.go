package common

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func MakeHttpGetResponseToByte(baseUrl string) (responseByte []byte, err error) {

	response, err := http.Get(baseUrl)

	responseByte, err = ioutil.ReadAll(response.Body)

	return
}

func MakeHttpPostResponseToByte(baseUrl string, requestByte []byte) (responseByte []byte, err error) {

	requestBody := bytes.NewBuffer(requestByte)

	response, err := http.Post(baseUrl, echo.MIMEApplicationJSON, requestBody)
	responseByte, err = ioutil.ReadAll(response.Body)

	return
}

func MakeApiUrl() {

}
