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

func OrderNoteRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewOrderNoteRepository(db, domain.CollectionOrderNote)
	lc := &controller.OrderNoteController{
		OrderNoteUsecase: usecase.NewOrderNoteUsecase(ur, timeout),
		Env:              env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/order_notes", lc.Fetch)
	itemGroup.GET("/order_note", lc.FetchByID)
	itemGroup.POST("/order_note", lc.Create)
	itemGroup.POST("/order_notes", lc.CreateMany)
	itemGroup.PUT("/order_note", lc.Update)
	itemGroup.DELETE("/order_note", lc.Delete)
}
