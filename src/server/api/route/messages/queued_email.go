package route

import (
	"time"

	controller "earnforglance/server/api/controller/messages"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/messages"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/messages"
	usecase "earnforglance/server/usecase/messages"

	"github.com/gin-gonic/gin"
)

func QueuedEmailRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewQueuedEmailRepository(db, domain.CollectionQueuedEmail)
	lc := &controller.QueuedEmailController{
		QueuedEmailUsecase: usecase.NewQueuedEmailUsecase(ur, timeout),
		Env:                env,
	}

	itemGroup := group.Group("/api/v1/messages")

	itemGroup.GET("/queued_emails", lc.Fetch)
	itemGroup.GET("/queued_email", lc.FetchByID)
	itemGroup.POST("/queued_email", lc.Create)
	itemGroup.POST("/queued_emails", lc.CreateMany)
	itemGroup.PUT("/queued_email", lc.Update)
	itemGroup.DELETE("/queued_email", lc.Delete)
}
