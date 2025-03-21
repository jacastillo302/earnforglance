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

func DiscountManufacturerMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountManufacturerMappingRepository(db, domain.CollectionDiscountManufacturerMapping)
	lc := &controller.DiscountManufacturerMappingController{
		DiscountManufacturerMappingUsecase: usecase.NewDiscountManufacturerMappingUsecase(ur, timeout),
		Env:                                env,
	}

	group.GET("/discount_manufacturer_mappings", lc.Fetch)
	group.GET("/discount_manufacturer_mapping", lc.FetchByID)
	group.POST("/discount_manufacturer_mapping", lc.Create)
	group.PUT("/discount_manufacturer_mapping", lc.Update)
	group.DELETE("/discount_manufacturer_mapping", lc.Delete)
}
