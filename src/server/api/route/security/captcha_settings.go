package route

import (
	"time"

	controller "earnforglance/server/api/controller/security"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/security"

	repository "earnforglance/server/repository/security"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/security"

	"github.com/gin-gonic/gin"
)

func CaptchaSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewCaptchaSettingsRepository(db, domain.CollectionCaptchaSettings)
	lc := &controller.CaptchaSettingsController{
		CaptchaSettingsUsecase: usecase.NewCaptchaSettingsUsecase(ur, timeout),
		Env:                    env,
	}
	itemGroup := group.Group("/api/v1/security")
	itemGroup.GET("/captcha_settings", lc.Fetch)
	itemGroup.GET("/captcha_setting", lc.FetchByID)
	itemGroup.POST("/captcha_setting", lc.Create)
	itemGroup.POST("/captcha_settings", lc.CreateMany)
	itemGroup.PUT("/captcha_setting", lc.Update)
	itemGroup.DELETE("/captcha_setting", lc.Delete)
}
