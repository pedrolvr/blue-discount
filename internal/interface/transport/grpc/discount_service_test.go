package grpc_test

import (
	usecase_mocks "blue-discount/internal/app/usecase/mocks"
	"blue-discount/internal/domain/model"
	infra_error "blue-discount/internal/infra/error"
	"blue-discount/internal/infra/repository"
	grpc_interface "blue-discount/internal/interface/transport/grpc"
	proto_v1 "blue-discount/proto/v1"
	"context"
	"errors"
	"log"
	"net"

	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	hostPort           = "localhost:7777"
	userID             = "ec1400b2-d1c3-471a-9fa3-c533a1ef1a3a"
	productID          = "15b69bd4-0357-45ed-96ca-62c308ed77e8"
	blackFridayPercent = int32(10)
	birthdayPercent    = int32(5)
	maxDiscount        = int32(10)
)

func newUser() model.User {
	ID, _ := uuid.FromString(userID)
	return model.NewUser(ID, nil)
}

func newProduct() model.Product {
	ID, _ := uuid.FromString(productID)
	return model.NewProduct(ID, 0)
}

func newDiscount(v int64, p int32) model.Discount {
	return model.NewDiscount(v, p)
}

func newPurchase(discount model.Discount) model.Purchase {
	return model.Purchase{
		User:     newUser(),
		Product:  newProduct(),
		Discount: discount,
	}
}

func dialer(registerService func(*grpc.Server)) func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	registerService(server)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func grpcClient(ctx context.Context, purchaseUsecaseMock *usecase_mocks.MockPurchaseUsecase) (*grpc.ClientConn, error) {
	register := func(server *grpc.Server) {
		proto_v1.RegisterDiscountServiceServer(server,
			grpc_interface.NewDiscountServer(purchaseUsecaseMock, infra_error.ErrorMiddleware()))
	}

	return grpc.DialContext(ctx, "", grpc.WithInsecure(),
		grpc.WithContextDialer(dialer(register)))
}

var _ = Describe("discount grpc service", func() {
	var (
		mockCtrl            *gomock.Controller
		purchaseUsecaseMock *usecase_mocks.MockPurchaseUsecase
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		purchaseUsecaseMock = usecase_mocks.NewMockPurchaseUsecase(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Discount()", func() {
		It("should get 450 and 10% of discount", func() {
			discount := newDiscount(450, 10)
			purchase := newPurchase(discount)

			purchaseUsecaseMock.
				EXPECT().
				Discount(userID, productID).
				Return(purchase, nil)

			ctx := context.Background()

			cc, err := grpcClient(ctx, purchaseUsecaseMock)

			Ω(err).ShouldNot(HaveOccurred())
			defer cc.Close()

			c := proto_v1.NewDiscountServiceClient(cc)

			r, err := c.Calculate(
				ctx,
				&proto_v1.DiscountRequest{
					UserId:    userID,
					ProductId: productID,
				},
			)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(r.UserId).Should(Equal(userID))
			Ω(r.ProductId).Should(Equal(productID))
			Ω(r.Discount.Value).Should(Equal(discount.Value))
			Ω(r.Discount.Percent).Should(Equal(discount.Percent))
		})

		It("should get error code 500", func() {
			purchaseUsecaseMock.
				EXPECT().
				Discount(userID, productID).
				Return(model.Purchase{}, errors.New("any"))

			ctx := context.Background()

			cc, err := grpcClient(ctx, purchaseUsecaseMock)

			Ω(err).ShouldNot(HaveOccurred())
			defer cc.Close()

			c := proto_v1.NewDiscountServiceClient(cc)

			_, err = c.Calculate(
				ctx,
				&proto_v1.DiscountRequest{
					UserId:    userID,
					ProductId: productID,
				},
			)

			Ω(err).Should(HaveOccurred())

			e, ok := status.FromError(err)

			Ω(ok).Should(Equal(true))
			Ω(e.Code()).Should(Equal(codes.Internal))
		})

		It("should get error code 404", func() {
			purchaseUsecaseMock.
				EXPECT().
				Discount(userID, productID).
				Return(model.Purchase{}, repository.ErrRowNotFound)

			ctx := context.Background()

			cc, err := grpcClient(ctx, purchaseUsecaseMock)

			Ω(err).ShouldNot(HaveOccurred())
			defer cc.Close()

			c := proto_v1.NewDiscountServiceClient(cc)

			_, err = c.Calculate(
				ctx,
				&proto_v1.DiscountRequest{
					UserId:    userID,
					ProductId: productID,
				},
			)

			Ω(err).Should(HaveOccurred())

			e, ok := status.FromError(err)

			Ω(ok).Should(Equal(true))
			Ω(e.Code()).Should(Equal(codes.NotFound))
		})
	})
})
