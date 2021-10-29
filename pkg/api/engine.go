package api

import (
	"coding-challenge-go/pkg/api/notifier"
	product_v1 "coding-challenge-go/pkg/api/product/v1"
	product_v2 "coding-challenge-go/pkg/api/product/v2"
	"coding-challenge-go/pkg/api/seller"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// CreateAPIEngine creates engine instance that serves API endpoints,
// consider it as a router for incoming requests.
func CreateAPIEngine(db *sql.DB) (*gin.Engine, error) {
	r := gin.New()
	sellerRepository := seller.NewRepository(db)
	notifier, err := notifier.NewNotifier()
	if err != nil {
		return nil, err
	}

	createAPIV1(r, db, sellerRepository, notifier)
	createAPIV2(r, db, sellerRepository, notifier)
	return r, nil
}

func createAPIV1(r *gin.Engine, db *sql.DB, sellerRepository seller.IRepository, noti notifier.INotifier) {
	productRepository := product_v1.NewRepository(db)
	productController := product_v1.NewControllerV1(productRepository, sellerRepository, noti)
	v1 := r.Group("api/v1")
	v1.GET("products", productController.List)
	v1.GET("product", productController.Get)
	v1.POST("product", productController.Post)
	v1.PUT("product", productController.Put)
	v1.DELETE("product", productController.Delete)

	sellerController := seller.NewController(sellerRepository)
	v1.GET("sellers", sellerController.List)
}

func createAPIV2(r *gin.Engine, db *sql.DB, sellerRepository seller.IRepository, noti notifier.INotifier) {
	productRepository := product_v2.NewRepository(db)
	productController := product_v2.NewControllerV2(productRepository, sellerRepository, noti)
	v2 := r.Group("api/v2")
	v2.GET("products", productController.List)
	v2.GET("product", productController.Get)
	v2.POST("product", productController.Post)
	v2.PUT("product", productController.Put)
	v2.DELETE("product", productController.Delete)

	sellerController := seller.NewController(sellerRepository)
	v2.GET("sellers", sellerController.List)
	v2.GET("sellers/top10", sellerController.GetTop10)
}
