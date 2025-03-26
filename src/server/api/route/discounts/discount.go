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

func DiscountRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountRepository(db, domain.CollectionDiscount)
	lc := &controller.DiscountController{
		DiscountUsecase: usecase.NewDiscountUsecase(ur, timeout),
		Env:             env,
	}

	itemGroup := group.Group("/api/v1/discounts")
	itemGroup.GET("/discounts", lc.Fetch)
	itemGroup.GET("/discount", lc.FetchByID)
	itemGroup.POST("/discount", lc.Create)
	itemGroup.POST("/discounts", lc.CreateMany)
	itemGroup.PUT("/discount", lc.Update)
	itemGroup.DELETE("/discount", lc.Delete)
}
