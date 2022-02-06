package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func SendResponse(w http.ResponseWriter, r *http.Request, responseCode int, responseData interface{}, errorMessage ErrorMessage) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var doc []byte
	var bodyLen = 0

	if responseCode == http.StatusOK {
		if responseData != nil {
			doc, _ = json.Marshal(responseData)
			bodyLen = len(doc)
		} else {
			w.Header().Set("Content-Length", "0")
		}
	} else {
		doc, _ = json.Marshal(errorMessage)
	}

	corrId := r.Header.Get("X-KM-Correlation-Id")
	caller := r.Header.Get("X-KM-CALLER")
	userId := r.Header.Get("X-Km-Userid")

	if caller == "" {
		caller = "CLIENT"
	}

	log.Printf("%s T[%s] U[%s] - [%s-RSP]	 %s (Len:%d)", LogLevel(), corrId, userId, caller, string(doc), bodyLen)

	if ResponseCompressEnable && bodyLen >= ResponseCompressMinSize {
		w.Header().Set("Content-Encoding", "application/gzip")
	}

	w.Header().Set("X-KM-TEMP", strconv.Itoa(responseCode))
	w.WriteHeader(responseCode)
	w.Write(doc)
}

func SendResponseWithoutLog(w http.ResponseWriter, r *http.Request, responseCode int, responseData interface{}, errorMessage ErrorMessage) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var doc []byte
	var bodyLen = 0

	if responseCode == http.StatusOK {
		if responseData != nil {
			doc, _ = json.Marshal(responseData)
			bodyLen = len(doc)
			//if ResponseCompressEnable && bodyLen > ResponseCompressMinSize {
			//w.Header().Set("Content-Encoding", "gzip")
			//}
		} else {
			w.Header().Set("Content-Length", "0")
		}
	} else {
		doc, _ = json.Marshal(errorMessage)
	}

	corrId := r.Header.Get("X-KM-Correlation-Id")
	caller := r.Header.Get("X-KM-CALLER")

	if caller == "" {
		caller = "CLIENT"
	}

	if strings.ToUpper(LogLevel()) == "DEBUG" {
		log.Printf("DEBUG T[%s] U[] - [%s-RSP]	%s (Len:%d)", corrId, caller, string(doc), bodyLen)
	} else {
		log.Printf("INFO T[%s] U[] - [%s-RSP] (Len:%d)", corrId, caller, bodyLen)
	}

	w.Header().Set("X-KM-TEMP", strconv.Itoa(responseCode))
	w.WriteHeader(responseCode)
	w.Write(doc)
}

func ReplaceMustacheWordAll(s string, mustacheWord string, new string) string {
	return strings.Replace(s, mustacheWord, new, -1)
}

func PrintJson(args ...interface{}) string {

	b, err := json.Marshal(args[0])

	if err != nil {
		fmt.Println("Json Marshal Error !")
		return ""
	}

	return string(b)
}

func AppendTid(tid string) (appendedString string) {
	b := make([]byte, 0, 6+len(tid))
	b = append(b, "T["...)
	b = append(b, tid...)
	b = append(b, "] - "...)
	return string(b)
}
