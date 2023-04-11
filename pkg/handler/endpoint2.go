package handler

import (
	"fmt"
	"net/http"
	"time"
)

var statusCode2 = http.StatusOK

func init() {
	go toggle2()
}

func EndpointHandler2(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("{ \"statusCode-process2\": %v  }", statusCode2)
	w.WriteHeader(statusCode2)
	w.Write([]byte(s))
}

func toggle2() {
	tick := time.Tick(time.Minute * 1)

	for {
		select {
		case <-tick:
			if statusCode2 == http.StatusOK {
				statusCode2 = http.StatusServiceUnavailable
			} else {
				statusCode2 = http.StatusOK
			}
		default:
		}
	}
}
