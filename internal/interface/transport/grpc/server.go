package grpc

import (
	"blue-discount/internal/app"
	"blue-discount/internal/app/usecase"
	"blue-discount/internal/infra/repository"
	proto_v1 "blue-discount/proto/v1"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func Start(db *gorm.DB, cfg app.Config) {
	port := cfg.Service.Port

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userRepo := repository.NewUserRepo(db)
	productRepo := repository.NewProductRepo(db)
	campaignRepo := repository.NewCampaignRepo(db)

	usc := usecase.NewPurchaseUsecaseImpl(userRepo, productRepo, campaignRepo, cfg.Discount)

	grpcServer := grpc.NewServer()

	proto_v1.RegisterDiscountServiceServer(grpcServer, NewDiscountServer(usc))

	fmt.Printf("Discount service is listening on port %d...\n", port)
	err = grpcServer.Serve(listener)
	fmt.Println("Serve() failed", err)
}
