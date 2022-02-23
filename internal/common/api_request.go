package common

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"time"
)

type ApiInfo struct {
	Tid       string
	UserId    string
	Component string
	Uri       string
	Method    string
}

var apiInfo = new(ApiInfo)

func SetApiInfo(context echo.Context) {

	apiInfo.Component = context.Request().Header.Get(echo.HeaderXKMCaller)
	if apiInfo.Component == "" {
		apiInfo.Component = "CLIENT"
	}

	apiInfo.Tid = makeTid(context)
	apiInfo.UserId = context.Request().Header.Get(echo.HeaderXKMUserId)
	apiInfo.Uri = context.Request().RequestURI
	apiInfo.Method = context.Request().Method
}

func GetApiInfo() *ApiInfo {
	return apiInfo
}

func makeTid(context echo.Context) (tid string) {

	nowSting := time.Now().Format("060102150405")
	randHex := fmt.Sprintf("%x", rand.Intn(100000000000000))[0:7]
	tid = nowSting + "-" + randHex

	context.Request().Header.Set(echo.HeaderXKMCorrelationId, tid)

	return
}

func setHeader(request *http.Request) {
	request.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	request.Header.Add(echo.HeaderXKMCorrelationId, apiInfo.Tid)
	request.Header.Add(echo.HeaderXKMUserId, apiInfo.UserId)
	request.Header.Add(echo.HeaderXKMCaller, apiInfo.Component)
}
