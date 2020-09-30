package grpc

import (
	"context"

	"blue-discount/internal/app/usecase"
	"blue-discount/internal/interface/dto"
	"blue-discount/internal/interface/facade"
	proto_v1 "blue-discount/proto/v1"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type DiscountHandler struct {
	calculate grpctransport.Handler
}

func (h *DiscountHandler) Calculate(ctx context.Context, r *proto_v1.DiscountRequest) (*proto_v1.DiscountResponse, error) {
	_, resp, err := h.calculate.ServeGRPC(ctx, r)

	if err != nil {
		return nil, err
	}

	return resp.(*proto_v1.DiscountResponse), nil
}

func decodeDiscountRequest(_ context.Context, r interface{}) (interface{}, error) {
	request := r.(*proto_v1.DiscountRequest)
	return dto.CalculateDiscountDTO{
		UserID:    request.UserId,
		ProductID: request.ProductId,
	}, nil
}

func encodeDiscountResponse(ctx context.Context, r interface{}) (interface{}, error) {
	d := r.(dto.PurchaseDTO)

	res := &proto_v1.DiscountResponse{
		UserId:    d.UserID,
		ProductId: d.ProductID,
	}

	res.Discount = &proto_v1.Discount{
		Value:   d.Discount.Value,
		Percent: d.Discount.Percent,
	}

	return res, nil
}

func NewDiscountServer(usc usecase.PurchaseUsecase, middlewares ...endpoint.Middleware) proto_v1.DiscountServiceServer {
	var endpoint endpoint.Endpoint

	endpoint = facade.MakeCalculateDiscountEndpoint(usc)

	for _, m := range middlewares {
		endpoint = m(endpoint)
	}

	return &DiscountHandler{
		calculate: grpctransport.NewServer(
			endpoint,
			decodeDiscountRequest,
			encodeDiscountResponse,
		),
	}
}
