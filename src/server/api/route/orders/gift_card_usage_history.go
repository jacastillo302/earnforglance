package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func GiftCardUsageHistoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGiftCardUsageHistoryRepository(db, domain.CollectionGiftCardUsageHistory)
	lc := &controller.GiftCardUsageHistoryController{
		GiftCardUsageHistoryUsecase: usecase.NewGiftCardUsageHistoryUsecase(ur, timeout),
		Env:                         env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/gift_card_usage_histories", lc.Fetch)
	itemGroup.GET("/gift_card_usage_history", lc.FetchByID)
	itemGroup.POST("/gift_card_usage_history", lc.Create)
	itemGroup.POST("/gift_card_usage_histories", lc.CreateMany)
	itemGroup.PUT("/gift_card_usage_history", lc.Update)
	itemGroup.DELETE("/gift_card_usage_history", lc.Delete)
}
