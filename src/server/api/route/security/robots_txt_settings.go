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

func RobotsTxtSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewRobotsTxtSettingsRepository(db, domain.CollectionRobotsTxtSettings)
	lc := &controller.RobotsTxtSettingsController{
		RobotsTxtSettingsUsecase: usecase.NewRobotsTxtSettingsUsecase(ur, timeout),
		Env:                      env,
	}
	itemGroup := group.Group("/api/v1/security")
	itemGroup.GET("/robots_txt_settings", lc.Fetch)
	itemGroup.GET("/robots_txt_setting", lc.FetchByID)
	itemGroup.POST("/robots_txt_setting", lc.Create)
	itemGroup.POST("/robots_txt_settings", lc.CreateMany)
	itemGroup.PUT("/robots_txt_setting", lc.Update)
	itemGroup.DELETE("/robots_txt_setting", lc.Delete)
}
