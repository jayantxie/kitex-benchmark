// Code generated by Kitex v0.4.2. DO NOT EDIT.
package echoserver

import (
	echo "github.com/cloudwego/kitex-benchmark/codec/thrift/kitex_gen/echo"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler echo.EchoServer, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
