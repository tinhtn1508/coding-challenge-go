package v1

import (
	"bytes"
	"coding-challenge-go/pkg/api/notifier"
	"coding-challenge-go/pkg/api/seller"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

func Test_controller_List(t *testing.T) {
	t.Run("failed to list product", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/v1/products", nil)

		repo := NewMockIRepository(mockCtrl)
		repo.EXPECT().list(0, 10).Return(nil, errors.New("failed"))

		seller := seller.NewMockIRepository(mockCtrl)
		provider := notifier.NewMockINotifier(mockCtrl)

		controller := &controller{
			repository:       repo,
			sellerRepository: seller,
			sellerNoti:       provider,
		}

		controller.List(c)
	})

	t.Run("successful to list product", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/v1/products", nil)

		repo := NewMockIRepository(mockCtrl)
		repo.EXPECT().list(0, 10).Return([]*product{}, nil)

		seller := seller.NewMockIRepository(mockCtrl)
		provider := notifier.NewMockINotifier(mockCtrl)

		controller := &controller{
			repository:       repo,
			sellerRepository: seller,
			sellerNoti:       provider,
		}

		controller.List(c)
	})
}

func Test_controller_Get(t *testing.T) {
	t.Run("failed to get product", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/v1/product?id=1", nil)

		repo := NewMockIRepository(mockCtrl)
		repo.EXPECT().findByUUID("1").Return(nil, errors.New("failed"))

		seller := seller.NewMockIRepository(mockCtrl)
		provider := notifier.NewMockINotifier(mockCtrl)

		controller := &controller{
			repository:       repo,
			sellerRepository: seller,
			sellerNoti:       provider,
		}

		controller.Get(c)
	})

	t.Run("successful to get product", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/v1/product?id=1", nil)

		repo := NewMockIRepository(mockCtrl)
		repo.EXPECT().findByUUID("1").Return(&product{}, nil)

		seller := seller.NewMockIRepository(mockCtrl)
		provider := notifier.NewMockINotifier(mockCtrl)

		controller := &controller{
			repository:       repo,
			sellerRepository: seller,
			sellerNoti:       provider,
		}

		controller.Get(c)
	})
}

func Test_controller_Post(t *testing.T) {
	t.Run("failed to post product", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("POST", "/v1/product", bytes.NewBuffer(
			[]byte(`{"name":"Berlin S.O.L.I.D. T-Shirt","brand":"Shirts Inc.","stock":150,"seller":"156c764b-f563-11e9-94e7-38baf859afa1"}`)))

		repo := NewMockIRepository(mockCtrl)
		repo.EXPECT().insert(gomock.Any()).Return(errors.New("failed"))

		sellerRepo := seller.NewMockIRepository(mockCtrl)

		sellerRepo.EXPECT().FindByUUID("156c764b-f563-11e9-94e7-38baf859afa1").Return(&seller.Seller{
			SellerID: 0,
			UUID:     "156c764b-f563-11e9-94e7-38baf859afa1",
			Name:     "test",
			Email:    "test@gmail.com",
			Phone:    "12345679",
		}, nil)

		provider := notifier.NewMockINotifier(mockCtrl)

		controller := &controller{
			repository:       repo,
			sellerRepository: sellerRepo,
			sellerNoti:       provider,
		}

		controller.Post(c)
	})

	t.Run("successful to post product", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("POST", "/v1/product", bytes.NewBuffer(
			[]byte(`{"name":"Berlin S.O.L.I.D. T-Shirt","brand":"Shirts Inc.","stock":150,"seller":"156c764b-f563-11e9-94e7-38baf859afa1"}`)))

		repo := NewMockIRepository(mockCtrl)
		repo.EXPECT().insert(gomock.Any()).Return(errors.New("failed"))

		sellerRepo := seller.NewMockIRepository(mockCtrl)

		sellerRepo.EXPECT().FindByUUID("156c764b-f563-11e9-94e7-38baf859afa1").Return(&seller.Seller{
			SellerID: 0,
			UUID:     "156c764b-f563-11e9-94e7-38baf859afa1",
			Name:     "test",
			Email:    "test@gmail.com",
			Phone:    "12345679",
		}, nil)

		provider := notifier.NewMockINotifier(mockCtrl)

		controller := &controller{
			repository:       repo,
			sellerRepository: sellerRepo,
			sellerNoti:       provider,
		}

		controller.Post(c)
	})
}
