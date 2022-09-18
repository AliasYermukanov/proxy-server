package main

import (
	"flag"
	"fmt"
	"github.com/AliasYermukanov/proxy-server/src/domain"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	liblogger "github.com/AliasYermukanov/proxy-server/src/util"
	"github.com/go-kit/kit/log/level"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/AliasYermukanov/proxy-server/src/api/middleware"
	transportHTTP "github.com/AliasYermukanov/proxy-server/src/api/transport/http"
	"github.com/AliasYermukanov/proxy-server/src/constructor"
	httpRPC "github.com/AliasYermukanov/proxy-server/src/rpc/http"
)

func main() {
	var container domain.Container
	container.Init()
	// parse flags
	httpPort := flag.String("http.port", ":8080", "HTTP listen address")
	flag.Parse()

	// init structured logger for the service
	logger := liblogger.NewServiceLogger("proxy-server")
	_ = level.Info(logger).Log("msg", "service started")

	httpClient := httpRPC.InitHTTPClient()
	externalRPC := httpRPC.NewRPC(httpClient, logger)
	proxySvc := constructor.NewProxyService(container, externalRPC, logger)

	proxyEndpoints := middleware.MakeEndpoints(proxySvc)

	// init HTTP handlers (transport layer)
	serverOptions := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(transportHTTP.EncodeError),
	}

	proxyHTTPHandler := transportHTTP.NewHTTPService(proxyEndpoints, serverOptions, logger)

	// add routes, prometheus and health check handlers
	http.Handle("/proxy-server/v1/proxy/", cors(proxyHTTPHandler))
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/check", healthCheck)

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

// cors
// headers append func
func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Lang")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

// healthCheck
// health check func
func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "ok")
}
