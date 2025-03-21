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

func OrderNoteRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewOrderNoteRepository(db, domain.CollectionOrderNote)
	lc := &controller.OrderNoteController{
		OrderNoteUsecase: usecase.NewOrderNoteUsecase(ur, timeout),
		Env:              env,
	}

	group.GET("/ordernotes", lc.Fetch)
	group.GET("/ordernote", lc.FetchByID)
	group.POST("/ordernote", lc.Create)
	group.PUT("/ordernote", lc.Update)
	group.DELETE("/ordernote", lc.Delete)
}
