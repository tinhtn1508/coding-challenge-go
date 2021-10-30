package seller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

func Test_controller_GetTop10(t *testing.T) {
	t.Run("failed to get top 10 product", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/v2/sellers/top10", nil)

		repo := NewMockIRepository(mockCtrl)
		repo.EXPECT().getTop(10).Return(nil, errors.New("failed"))

		controller := &controller{
			repository: repo,
		}

		controller.GetTop10(c)
	})

	t.Run("successful to get top 10 product", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request, _ = http.NewRequest("GET", "/v2/sellers/top10", nil)

		repo := NewMockIRepository(mockCtrl)
		repo.EXPECT().getTop(10).Return([]*Seller{}, nil)

		controller := &controller{
			repository: repo,
		}

		controller.GetTop10(c)
	})
}
