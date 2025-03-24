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

func DiscountCategoryMappingRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountCategoryMappingRepository(db, domain.CollectionDiscountCategoryMapping)
	lc := &controller.DiscountCategoryMappingController{
		DiscountCategoryMappingUsecase: usecase.NewDiscountCategoryMappingUsecase(ur, timeout),
		Env:                            env,
	}

	group.GET("/discount_category_mappings", lc.Fetch)
	group.GET("/discount_category_mapping", lc.FetchByID)
	group.POST("/discount_category_mapping", lc.Create)
	group.POST("/discount_category_mappings", lc.CreateMany)
	group.PUT("/discount_category_mapping", lc.Update)
	group.DELETE("/discount_category_mapping", lc.Delete)
}
