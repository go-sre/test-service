package host

import (
	"github.com/go-sre/core/exchange"
	"github.com/go-sre/core/runtime"
	"github.com/go-sre/host/accesslog"
	middleware2 "github.com/go-sre/host/middleware"
	"github.com/go-sre/test-service/pkg/resource"
	"net/http"
)

const (
	egressLogOperatorNameFmt  = "fs/egress_logging_operators.json"
	ingressLogOperatorNameFmt = "fs/ingress_logging_operators.json"
)

func Startup[E runtime.ErrorHandler, O runtime.OutputHandler](mux *http.ServeMux) (http.Handler, *runtime.Status) {
	var e E
	err := initLogging()
	if err != nil {
		return nil, e.Handle(nil, "/host/startup/logging", err)
	}
	initMux(mux)
	middleware2.AccessLogWrapTransport(exchange.Client)
	return middleware2.AccessHttpHostMetricsHandler(mux, ""), runtime.NewStatusOK()
}

func Shutdown() {
}

func initLogging() error {
	// Options that are defaulted to true for the statuses
	accesslog.SetIngressLogStatus(true)
	accesslog.SetEgressLogStatus(true)
	accesslog.SetPingLogStatus(true)

	// Enable logging function for access events middleware
	// middleware.SetLogFn(func(entry *data.Entry) {
	// log.Write[log.DebugOutputHandler, data.JsonFormatter](entry)
	//},
	//)
	//controller.SetLogFn(func(traffic string, start time.Time, duration time.Duration, req *http.Request, resp *http.Response, routeName string, timeout int, limit rate.Limit, burst int, retry, proxy, statusFlags string) {
	//	entry := accessdata.NewEntry(traffic, start, duration, req, resp, routeName, timeout, limit, burst, retry, proxy, statusFlags)
	//	accesslog.Write[accesslog.DebugOutputHandler, accessdata.JsonFormatter](entry)
	//},
	//)

	err := accesslog.CreateIngressOperators(func() ([]byte, error) {
		return resource.ReadFile(ingressLogOperatorNameFmt)
	})
	if err == nil {
		err = accesslog.CreateEgressOperators(func() ([]byte, error) {
			return resource.ReadFile(egressLogOperatorNameFmt)
		})
	}
	return err
}
