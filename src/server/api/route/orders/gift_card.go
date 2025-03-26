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

func GiftCardRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGiftCardRepository(db, domain.CollectionGiftCard)
	lc := &controller.GiftCardController{
		GiftCardUsecase: usecase.NewGiftCardUsecase(ur, timeout),
		Env:             env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/gift_cards", lc.Fetch)
	itemGroup.GET("/gift_card", lc.FetchByID)
	itemGroup.POST("/gift_card", lc.Create)
	itemGroup.POST("/gift_cards", lc.CreateMany)
	itemGroup.PUT("/gift_card", lc.Update)
	itemGroup.DELETE("/gift_card", lc.Delete)
}
