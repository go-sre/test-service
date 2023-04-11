package host

import (
	"github.com/go-sre/test-service/pkg/handler"
	"net/http"
	"net/http/pprof"
)

const (
	processPattern        = "/process"
	processPattern2       = "/process2"
	healthLivenessPattern = "/health/liveness"

	IndexPattern   = "/debug/pprof/"
	CmdLinePattern = "/debug/pprof/cmdline"
	ProfilePattern = "/debug/pprof/profile" // ?seconds=30
	SymbolPattern  = "/debug/pprof/symbol"
	TracePattern   = "/debug/pprof/trace"
)

func initMux(r *http.ServeMux) {

	addProfileRoutes(r)

	r.Handle(healthLivenessPattern, http.HandlerFunc(handler.HealthLivenessHandler))
	r.Handle(processPattern, http.HandlerFunc(handler.EndpointHandler))
	r.Handle(processPattern2, http.HandlerFunc(handler.EndpointHandler2))

}

func addProfileRoutes(r *http.ServeMux) {
	r.Handle(IndexPattern, http.HandlerFunc(pprof.Index))
	r.Handle(CmdLinePattern, http.HandlerFunc(pprof.Cmdline))
	r.Handle(ProfilePattern, http.HandlerFunc(pprof.Profile))
	r.Handle(SymbolPattern, http.HandlerFunc(pprof.Symbol))
	r.Handle(TracePattern, http.HandlerFunc(pprof.Trace))

}
