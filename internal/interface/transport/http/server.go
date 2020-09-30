package http

import (
	"blue-discount/internal/app"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Start(g *run.Group, logger log.Logger, cfg app.Config) {
	port := cfg.Service.HTTPPort

	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		level.Error(logger).Log("msg",
			fmt.Sprintf("listen http error: %v", err))
		os.Exit(1)
	}

	g.Add(func() error {
		level.Info(logger).Log("msg",
			fmt.Sprintf("HTTP server listening on %d", port))
		return http.Serve(l, http.DefaultServeMux)
	}, func(error) {
		l.Close()
	})
}
