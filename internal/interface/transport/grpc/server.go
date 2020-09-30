package grpc

import (
	"blue-discount/internal/app"
	"blue-discount/internal/app/usecase"
	infra_error "blue-discount/internal/infra/error"
	"blue-discount/internal/infra/log"
	"blue-discount/internal/infra/metric"
	"blue-discount/internal/infra/repository"
	proto_v1 "blue-discount/proto/v1"
	"fmt"
	"net"
	"os"

	kit_log "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func Start(g *run.Group, logger kit_log.Logger, db *gorm.DB, cfg app.Config) {
	port := cfg.Service.GRPCPort

	userRepo := repository.NewUserRepo(db)
	productRepo := repository.NewProductRepo(db)
	campaignRepo := repository.NewCampaignRepo(db)

	purchaseUsecase := usecase.NewPurchaseUsecaseImpl(userRepo,
		productRepo, campaignRepo, cfg.Discount)

	grpcServer := grpc.NewServer()

	discountServer := NewDiscountServer(
		purchaseUsecase,
		log.LoggingMiddleware(logger),
		infra_error.ErrorMiddleware(),
		metric.RequestInstrumenting("discount", "calculate"),
	)

	proto_v1.RegisterDiscountServiceServer(grpcServer, discountServer)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		level.Error(logger).Log("msg",
			fmt.Sprintf("listen grpc error: %v", err))
		os.Exit(1)
	}

	g.Add(func() error {
		level.Info(logger).Log("msg",
			fmt.Sprintf("GRPC server listening on %d", port))
		return grpcServer.Serve(l)
	}, func(error) {
		l.Close()
	})
}
