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

func EmailAccountSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewEmailAccountSettingsRepository(db, domain.CollectionEmailAccountSettings)
	lc := &controller.EmailAccountSettingsController{
		EmailAccountSettingsUsecase: usecase.NewEmailAccountSettingsUsecase(ur, timeout),
		Env:                         env,
	}
	itemGroup := group.Group("/api/v1/messages")
	itemGroup.GET("/email_account_settings", lc.Fetch)
	itemGroup.GET("/email_account_setting", lc.FetchByID)
	itemGroup.POST("/email_account_setting", lc.Create)
	itemGroup.POST("/email_account_settings", lc.CreateMany)
	itemGroup.PUT("/email_account_setting", lc.Update)
	itemGroup.DELETE("/email_account_setting", lc.Delete)
}
