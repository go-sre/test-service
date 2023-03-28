package handler

import (
	"fmt"
	"net/http"
	"time"
)

var statusCode = http.StatusOK

func init() {
	go toggle()
}

func EndpointHandler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("{ \"statusCode\": %v  }", statusCode)
	w.WriteHeader(statusCode)
	w.Write([]byte(s))
}

func toggle() {
	tick := time.Tick(time.Minute * 1)

	for {
		select {
		case <-tick:
			if statusCode == http.StatusOK {
				statusCode = http.StatusServiceUnavailable
			} else {
				statusCode = http.StatusOK
			}
		default:
		}
	}
}
