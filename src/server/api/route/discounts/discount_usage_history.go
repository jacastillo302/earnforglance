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

	group.GET("/discount_usage_histories", lc.Fetch)
	group.GET("/discount_usage_history", lc.FetchByID)
	group.POST("/discount_usage_history", lc.Create)
	group.PUT("/discount_usage_history", lc.Update)
	group.DELETE("/discount_usage_history", lc.Delete)
}
