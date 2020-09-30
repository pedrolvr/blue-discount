package metric

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var labelNames = []string{"method", "error"}

func newCounter(service, name string, labelNames []string) *kitprometheus.Counter {
	return kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "",
		Subsystem: service,
		Name:      name + "_count",
		Help:      "count # of requests",
	}, labelNames)
}

func newSummary(service, name string, labelNames []string) *kitprometheus.Summary {
	return kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "",
		Subsystem: service,
		Name:      name + "_summary",
		Help:      "request summary in milliseconds",
	}, labelNames)
}

func RequestInstrumenting(service, method string) endpoint.Middleware {
	requestCount := newCounter(service, method, labelNames)
	requestLatency := newSummary(service, method, labelNames)

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				lvs := []string{"method", method, "error", fmt.Sprint(err != nil)}

				requestCount.With(lvs...).Add(1)
				requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
			}(time.Now())

			response, err = next(ctx, request)

			return
		}
	}
}
