package route

import (
	"time"

	controller "earnforglance/server/api/controller/discounts"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/discounts"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/discounts"
	usecase "earnforglance/server/usecase/discounts"

	"github.com/gin-gonic/gin"
)

func DiscountProductMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountProductMappingRepository(db, domain.CollectionDiscountProductMapping)
	lc := &controller.DiscountProductMappingController{
		DiscountProductMappingUsecase: usecase.NewDiscountProductMappingUsecase(ur, timeout),
		Env:                           env,
	}

	group.GET("/discount_product_mappings", lc.Fetch)
	group.GET("/discount_product_mapping", lc.FetchByID)
	group.POST("/discount_product_mapping", lc.Create)
	group.POST("/discount_product_mappings", lc.CreateMany)
	group.PUT("/discount_product_mapping", lc.Update)
	group.DELETE("/discount_product_mapping", lc.Delete)
}
