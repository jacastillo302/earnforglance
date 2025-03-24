package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/orders"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func GiftCardUsageHistoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGiftCardUsageHistoryRepository(db, domain.CollectionGiftCardUsageHistory)
	lc := &controller.GiftCardUsageHistoryController{
		GiftCardUsageHistoryUsecase: usecase.NewGiftCardUsageHistoryUsecase(ur, timeout),
		Env:                         env,
	}

	group.GET("/gift_card_usage_histories", lc.Fetch)
	group.GET("/gift_card_usage_history", lc.FetchByID)
	group.POST("/gift_card_usage_history", lc.Create)
	group.POST("/gift_card_usage_histories", lc.CreateMany)
	group.PUT("/gift_card_usage_history", lc.Update)
	group.DELETE("/gift_card_usage_history", lc.Delete)
}
