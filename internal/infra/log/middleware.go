package log

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func LoggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var rawError error

			defer func(begin time.Time) {
				if rawError != nil {
					level.Error(logger).Log(
						"msg", rawError,
						"request", request,
						"response", response,
						"took", time.Since(begin),
					)
				} else {
					level.Info(logger).Log(
						"request", request,
						"response", response,
						"took", time.Since(begin),
					)
				}
			}(time.Now())

			response, err = next(ctx, request)
			return
		}
	}
}
