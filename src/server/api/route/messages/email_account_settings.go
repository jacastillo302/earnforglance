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

func EmailAccountSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewEmailAccountSettingsRepository(db, domain.CollectionEmailAccountSettings)
	lc := &controller.EmailAccountSettingsController{
		EmailAccountSettingsUsecase: usecase.NewEmailAccountSettingsUsecase(ur, timeout),
		Env:                         env,
	}

	group.GET("/email_account_settings", lc.Fetch)
	group.GET("/email_account_setting", lc.FetchByID)
	group.POST("/email_account_setting", lc.Create)
	group.POST("/email_account_settings", lc.CreateMany)
	group.PUT("/email_account_setting", lc.Update)
	group.DELETE("/email_account_setting", lc.Delete)
}
