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

func DiscountManufacturerMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountManufacturerMappingRepository(db, domain.CollectionDiscountManufacturerMapping)
	lc := &controller.DiscountManufacturerMappingController{
		DiscountManufacturerMappingUsecase: usecase.NewDiscountManufacturerMappingUsecase(ur, timeout),
		Env:                                env,
	}

	itemGroup := group.Group("/api/v1/discounts")
	itemGroup.GET("/discount_manufacturer_mappings", lc.Fetch)
	itemGroup.GET("/discount_manufacturer_mapping", lc.FetchByID)
	itemGroup.POST("/discount_manufacturer_mapping", lc.Create)
	itemGroup.POST("/discount_manufacturer_mappings", lc.CreateMany)
	itemGroup.PUT("/discount_manufacturer_mapping", lc.Update)
	itemGroup.DELETE("/discount_manufacturer_mapping", lc.Delete)
}
