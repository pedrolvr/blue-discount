package grpc

import (
	"context"
	"errors"

	"blue-discount/internal/app/usecase"
	"blue-discount/internal/infra/repository"
	"blue-discount/internal/interface/dto"
	"blue-discount/internal/interface/facade"
	proto_v1 "blue-discount/proto/v1"

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

	if d.Err != nil {
		res.Err = errorHandler(d.Err)
		return res, nil
	}

	res.Discount = &proto_v1.Discount{
		Value:   d.Discount.Value,
		Percent: d.Discount.Percent,
	}

	return res, nil
}

func errorHandler(err error) *proto_v1.ErrorInfo {
	errCode := 500
	reason := "internal error"

	if errors.Is(err, repository.ErrRowNotFound) {
		errCode = 404
		reason = err.Error()
	}

	return &proto_v1.ErrorInfo{
		Reason: reason,
		Code:   int32(errCode),
	}
}

func NewDiscountServer(usc usecase.PurchaseUsecase) proto_v1.DiscountServiceServer {
	return &DiscountHandler{
		calculate: grpctransport.NewServer(
			facade.MakeCalculateDiscountEndpoint(usc),
			decodeDiscountRequest,
			encodeDiscountResponse,
		),
	}
}
