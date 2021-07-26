// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/seniorGolang/tg/example/interfaces"
)

type serverExampleRPC struct {
	svc  interfaces.ExampleRPC
	test ExampleRPCTest
}

type MiddlewareSetExampleRPC interface {
	Wrap(m MiddlewareExampleRPC)
	WrapTest(m MiddlewareExampleRPCTest)

	WithTrace()
	WithMetrics()
	WithLog(log zerolog.Logger)
}

func newServerExampleRPC(svc interfaces.ExampleRPC) *serverExampleRPC {
	return &serverExampleRPC{
		svc:  svc,
		test: svc.Test,
	}
}

func (srv *serverExampleRPC) Wrap(m MiddlewareExampleRPC) {
	srv.svc = m(srv.svc)
	srv.test = srv.svc.Test
}

func (srv *serverExampleRPC) Test(ctx context.Context, arg0 int, arg1 string, opts ...interface{}) (ret1 int, ret2 string, err error) {
	return srv.test(ctx, arg0, arg1, opts...)
}

func (srv *serverExampleRPC) WrapTest(m MiddlewareExampleRPCTest) {
	srv.test = m(srv.test)
}

func (srv *serverExampleRPC) WithTrace() {
	srv.Wrap(traceMiddlewareExampleRPC)
}

func (srv *serverExampleRPC) WithMetrics() {
	srv.Wrap(metricsMiddlewareExampleRPC)
}

func (srv *serverExampleRPC) WithLog(log zerolog.Logger) {
	srv.Wrap(loggerMiddlewareExampleRPC(log))
}