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
	itemGroup.GET("/ordernotes", lc.Fetch)
	itemGroup.GET("/ordernote", lc.FetchByID)
	itemGroup.POST("/ordernote", lc.Create)
	itemGroup.POST("/ordernotes", lc.CreateMany)
	itemGroup.PUT("/ordernote", lc.Update)
	itemGroup.DELETE("/ordernote", lc.Delete)
}
