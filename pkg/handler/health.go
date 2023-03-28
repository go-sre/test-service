package handler

import (
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"net/http"
)

func HealthLivenessHandler(w http.ResponseWriter, r *http.Request) {
	var status = runtime.NewStatusOK()
	if status.OK() {
		exchange.WriteResponse(w, []byte("up"), status)
	} else {
		exchange.WriteResponse(w, nil, status)
	}
}
