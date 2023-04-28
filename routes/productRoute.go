package routes

import (
	"e-commerce-api/database"
	productHandler "e-commerce-api/feature/product/delivery/http"
	productRepository "e-commerce-api/feature/product/repository"
	productUsecase "e-commerce-api/feature/product/usecase"

	"github.com/gin-gonic/gin"
)

func ProductRoute(v1 *gin.RouterGroup) {

	shopRoute := v1.Group("/products")

	newProductRepository := productRepository.NewProductRepository(database.DB)
	newProductUsecase := productUsecase.NewProductUsecase(newProductRepository)
	newProductHandler := productHandler.NewProductHandler(newProductUsecase)

	shopRoute.GET("/all", newProductHandler.GetAllProducts)
	shopRoute.GET("/all/:shopId", newProductHandler.GetAllProductsByShopId)
	shopRoute.GET("/:productId", newProductHandler.GetProductById)
	shopRoute.GET("/filter/:categoryId", newProductHandler.FilterProductByCategoryId)
	shopRoute.GET("/filter/:categoryId/:shopId", newProductHandler.FilterProductByCategoryIdAndShopId)
}
