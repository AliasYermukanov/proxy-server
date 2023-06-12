package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log/level"

	"github.com/AliasYermukanov/proxy-server/src/api/middleware"
	"github.com/AliasYermukanov/proxy-server/src/api/service"
	transportHTTP "github.com/AliasYermukanov/proxy-server/src/api/transport/http"
	inmemRepo "github.com/AliasYermukanov/proxy-server/src/repository/inmem"
	httpRPC "github.com/AliasYermukanov/proxy-server/src/rpc/http"
	"github.com/AliasYermukanov/proxy-server/src/util"
)

func main() {
	ctx := context.Background()
	// parse flags
	httpPort := flag.String("http.port", ":8080", "HTTP listen address")
	flag.Parse()

	// init structured logger for the service
	logger := util.NewServiceLogger("proxy-server")
	_ = level.Info(logger).Log("msg", "service started")

	inMemCacheRepo, err := inmemRepo.NewStore(ctx, logger)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	httpClient := httpRPC.InitHTTPClient()
	externalRPC := httpRPC.NewRPC(httpClient, logger)
	proxySvc := service.NewService(inMemCacheRepo, externalRPC, logger)

	proxyEndpoints := middleware.MakeEndpoints(proxySvc)

	// init HTTP handlers (transport layer)
	serverOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(transportHTTP.EncodeError),
	}

	proxyHTTPHandler := transportHTTP.NewHTTPService(proxyEndpoints, serverOptions, logger)

	// add routes, prometheus and health check handlers
	http.Handle("/proxy-server/v1/proxy/", util.Cors(proxyHTTPHandler))
	http.HandleFunc("/check", util.HealthCheck)

	// init errors chan
	errs := make(chan error)

	// make chan for syscall
	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	// init HTTP server
	go func() {
		_ = level.Info(logger).Log("transport", "HTTP", "port", *httpPort)
		errs <- http.ListenAndServe(*httpPort, nil)
	}()

	defer func() {
		_ = level.Info(logger).Log("msg", "service ended")
	}()

	_ = level.Error(logger).Log("exit", <-errs)
}
