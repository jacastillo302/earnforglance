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

func RobotsTxtSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewRobotsTxtSettingsRepository(db, domain.CollectionRobotsTxtSettings)
	lc := &controller.RobotsTxtSettingsController{
		RobotsTxtSettingsUsecase: usecase.NewRobotsTxtSettingsUsecase(ur, timeout),
		Env:                      env,
	}

	group.GET("/robots_txt_settings", lc.Fetch)
	group.GET("/robots_txt_setting", lc.FetchByID)
	group.POST("/robots_txt_setting", lc.Create)
	group.POST("/robots_txt_settings", lc.CreateMany)
	group.PUT("/robots_txt_setting", lc.Update)
	group.DELETE("/robots_txt_setting", lc.Delete)
}
