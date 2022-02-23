package common

import (
	"log"
	"time"
)

func RequestLog(componentCheck string, methodCheck string, uriCheck string, requestByte string) time.Time {

	component, method, uri := checkApi(componentCheck, methodCheck, uriCheck)

	log.Printf("INFO T[%s] U[%s] - [%s-REQ] %s %s %s\n",
		GetApiInfo().Tid,
		GetApiInfo().UserId,
		component,
		method,
		uri,
		requestByte)

	return time.Now()
}

// TODO status code adding
func ResponseLog(componentCheck string, methodCheck string, uriCheck string, v interface{}, latency int64) {

	component, method, uri := checkApi(componentCheck, methodCheck, uriCheck)
	log.Printf("INFO T[%s] U[%s] - [%s-RES] %s %s %s %dms\n",
		GetApiInfo().Tid,
		GetApiInfo().UserId,
		component,
		method,
		uri,
		v,
		latency)
}

func checkApi(componentCheck string, methodCheck string, uriCheck string) (string, string, string) {
	component, method, uri := componentCheck, methodCheck, uriCheck
	if componentCheck == "" {
		component = GetApiInfo().Component
	}
	if methodCheck == "" {
		method = GetApiInfo().Method
	}
	if uriCheck == "" {
		uri = GetApiInfo().Uri
	}
	return component, method, uri
}

func Debug(v ...interface{}) {
	log.Printf("DEBUG T[%s] U[%s] %s\n",
		GetApiInfo().Tid,
		GetApiInfo().UserId,
		v)
}

func Info(v ...interface{}) {
	log.Printf("INFO T[%s] U[%s] %s\n",
		GetApiInfo().Tid,
		GetApiInfo().UserId,
		v)
}

func Warn(v ...interface{}) {
	log.Printf("WARN T[%s] U[%s] %s\n",
		GetApiInfo().Tid,
		GetApiInfo().UserId,
		v)
}

func Error(v ...interface{}) {
	log.Printf("ERROR T[%s] U[%s] %s\n",
		GetApiInfo().Tid,
		GetApiInfo().UserId,
		v)
}

func Fatal(v ...interface{}) {
	log.Printf("FATAL T[%s] U[%s] %s\n",
		GetApiInfo().Tid,
		GetApiInfo().UserId,
		v)
}
