package route

import (
	"time"

	controller "earnforglance/server/api/controller/discounts"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/discounts"

	repository "earnforglance/server/repository/discounts"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/discounts"

	"github.com/gin-gonic/gin"
)

func DiscountProductMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountProductMappingRepository(db, domain.CollectionDiscountProductMapping)
	lc := &controller.DiscountProductMappingController{
		DiscountProductMappingUsecase: usecase.NewDiscountProductMappingUsecase(ur, timeout),
		Env:                           env,
	}

	itemGroup := group.Group("/api/v1/discounts")
	itemGroup.GET("/discount_product_mappings", lc.Fetch)
	itemGroup.GET("/discount_product_mapping", lc.FetchByID)
	itemGroup.POST("/discount_product_mapping", lc.Create)
	itemGroup.POST("/discount_product_mappings", lc.CreateMany)
	itemGroup.PUT("/discount_product_mapping", lc.Update)
	itemGroup.DELETE("/discount_product_mapping", lc.Delete)
}
