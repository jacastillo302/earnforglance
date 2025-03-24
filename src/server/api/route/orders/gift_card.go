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

func GiftCardRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewGiftCardRepository(db, domain.CollectionGiftCard)
	lc := &controller.GiftCardController{
		GiftCardUsecase: usecase.NewGiftCardUsecase(ur, timeout),
		Env:             env,
	}

	group.GET("/gift_cards", lc.Fetch)
	group.GET("/gift_card", lc.FetchByID)
	group.POST("/gift_card", lc.Create)
	group.POST("/gift_cards", lc.CreateMany)
	group.PUT("/gift_card", lc.Update)
	group.DELETE("/gift_card", lc.Delete)
}
