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

func DiscountUsageHistoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewDiscountUsageHistoryRepository(db, domain.CollectionDiscountUsageHistory)
	lc := &controller.DiscountUsageHistoryController{
		DiscountUsageHistoryUsecase: usecase.NewDiscountUsageHistoryUsecase(ur, timeout),
		Env:                         env,
	}

	itemGroup := group.Group("/api/v1/discounts")
	itemGroup.GET("/discount_usage_histories", lc.Fetch)
	itemGroup.GET("/discount_usage_history", lc.FetchByID)
	itemGroup.POST("/discount_usage_history", lc.Create)
	itemGroup.POST("/discount_usage_histories", lc.CreateMany)
	itemGroup.PUT("/discount_usage_history", lc.Update)
	itemGroup.DELETE("/discount_usage_history", lc.Delete)
}
