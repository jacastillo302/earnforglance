package route

import (
	"time"

	controller "earnforglance/server/api/controller/messages"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/messages"

	repository "earnforglance/server/repository/messages"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/messages"

	"github.com/gin-gonic/gin"
)

func EmailAccountRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewEmailAccountRepository(db, domain.CollectionEmailAccount)
	lc := &controller.EmailAccountController{
		EmailAccountUsecase: usecase.NewEmailAccountUsecase(ur, timeout),
		Env:                 env,
	}
	itemGroup := group.Group("/api/v1/messages")
	itemGroup.GET("/email_accounts", lc.Fetch)
	itemGroup.GET("/email_account", lc.FetchByID)
	itemGroup.POST("/email_account", lc.Create)
	itemGroup.POST("/email_accounts", lc.CreateMany)
	itemGroup.PUT("/email_account", lc.Update)
	itemGroup.DELETE("/email_account", lc.Delete)
}
