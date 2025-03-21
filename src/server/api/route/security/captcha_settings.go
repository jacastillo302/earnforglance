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

	group.GET("/captcha_settingss", lc.Fetch)
	group.GET("/captcha_settings", lc.FetchByID)
	group.POST("/captcha_settings", lc.Create)
	group.PUT("/captcha_settings", lc.Update)
	group.DELETE("captcha_settings", lc.Delete)
}
