package error

import (
	"blue-discount/internal/infra/repository"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = next(ctx, request)

			if err != nil {
				err = errorHandler(err)
			}

			return
		}
	}
}

func errorHandler(err error) error {
	c := codes.Internal
	reason := "internal error"

	if errors.Is(err, repository.ErrRowNotFound) {
		c = codes.NotFound
		reason = err.Error()
	}

	return status.Error(c, reason)
}
