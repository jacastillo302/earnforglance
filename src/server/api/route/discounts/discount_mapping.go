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

func DiscountMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountMappingRepository(db, domain.CollectionDiscountMapping)
	lc := &controller.DiscountMappingController{
		DiscountMappingUsecase: usecase.NewDiscountMappingUsecase(ur, timeout),
		Env:                    env,
	}

	itemGroup := group.Group("/api/v1/discounts")
	itemGroup.GET("/discount_mappings", lc.Fetch)
	itemGroup.GET("/discount_mapping", lc.FetchByID)
	itemGroup.POST("/discount_mapping", lc.Create)
	itemGroup.POST("/discount_mappings", lc.CreateMany)
	itemGroup.PUT("/discount_mapping", lc.Update)
	itemGroup.DELETE("/discount_mapping", lc.Delete)
}
