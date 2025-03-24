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

func EmailAccountRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewEmailAccountRepository(db, domain.CollectionEmailAccount)
	lc := &controller.EmailAccountController{
		EmailAccountUsecase: usecase.NewEmailAccountUsecase(ur, timeout),
		Env:                 env,
	}

	group.GET("/email_accounts", lc.Fetch)
	group.GET("/email_account", lc.FetchByID)
	group.POST("/email_account", lc.Create)
	group.POST("/email_accounts", lc.CreateMany)
	group.PUT("/email_account", lc.Update)
	group.DELETE("/email_account", lc.Delete)
}
