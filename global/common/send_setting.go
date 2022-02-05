package common

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

package common

import (
"compress/gzip"
"fmt"
"github.com/json-iterator/go"
"io"
"log"
"net/http"
"strconv"
"strings"
"sync"
)

func SendResponse(w http.ResponseWriter, r *http.Request, responseCode int, responseData interface{}, errorMessage ErrorMessage) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var doc []byte
	var bodyLen = 0

	if responseCode == http.StatusOK {
		if responseData != nil {
			doc, _ = jsoniter.Marshal(responseData)
			bodyLen = len(doc)

			//if ResponseCompressEnable && bodyLen >= ResponseCompressMinSize {
			//w.Header().Set("Content-Encoding", "gzip")
			//}

		} else {
			w.Header().Set("Content-Length", "0")
		}
	} else {
		doc, _ = jsoniter.Marshal(errorMessage)
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
			doc, _ = jsoniter.Marshal(responseData)
			bodyLen = len(doc)
			//if ResponseCompressEnable && bodyLen > ResponseCompressMinSize {
			//w.Header().Set("Content-Encoding", "gzip")
			//}
		} else {
			w.Header().Set("Content-Length", "0")
		}
	} else {
		doc, _ = jsoniter.Marshal(errorMessage)
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

// TODO : GZIP

// Create a Pool that contains previously used Writers and
// can create new ones if we run out.
var Zippers = sync.Pool{New: func() interface{} {
	return gzip.NewWriter(nil)
}}

// Gzip Compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {

	if len(b) < ResponseCompressMinSize {
		return w.ResponseWriter.Write(b)
	}

	return w.Writer.Write(b)
}

// Wrap a http.Handler to support transparent gzip encoding.
func GzipHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if ResponseCompressEnable {
			// Do not set encoding here
			//w.Header().Set("Content-Encoding", "gzip")

			// Get a Writer from the Pool
			gz := Zippers.Get().(*gzip.Writer)

			// When done, put the Writer back in to the Pool
			defer Zippers.Put(gz)

			// We use Reset to set the writer we want to use.
			gz.Reset(w)
			defer gz.Close()

			gzw := &gzipResponseWriter{Writer: gz, ResponseWriter: w}
			handler.ServeHTTP(gzw, r)
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}

func ReplaceMustacheWordAll(s string, mustacheWord string, new string) string {
	return strings.Replace(s, mustacheWord, new, -1)
}

func PrintJson(args ...interface{}) string {

	b, err := jsoniter.Marshal(args[0])

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
