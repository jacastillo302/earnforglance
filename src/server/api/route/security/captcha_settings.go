package route

import (
	"time"

	controller "earnforglance/server/api/controller/security"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/security"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/security"
	usecase "earnforglance/server/usecase/security"

	"github.com/gin-gonic/gin"
)

func CaptchaSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCaptchaSettingsRepository(db, domain.CollectionCaptchaSettings)
	lc := &controller.CaptchaSettingsController{
		CaptchaSettingsUsecase: usecase.NewCaptchaSettingsUsecase(ur, timeout),
		Env:                    env,
	}

	group.GET("/captcha_settings", lc.Fetch)
	group.GET("/captcha_setting", lc.FetchByID)
	group.POST("/captcha_setting", lc.Create)
	group.POST("/captcha_settings", lc.CreateMany)
	group.PUT("/captcha_setting", lc.Update)
	group.DELETE("/captcha_setting", lc.Delete)
}
