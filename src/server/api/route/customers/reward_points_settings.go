package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	repository "earnforglance/server/repository/customers"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func RewardPointsSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewRewardPointsSettingsRepository(db, domain.CollectionRewardPointsSettings)
	lc := &controller.RewardPointsSettingsController{
		RewardPointsSettingsUsecase: usecase.NewRewardPointsSettingsUsecase(ur, timeout),
		Env:                         env,
	}

	itemGroup := group.Group("/api/v1/customers")
	itemGroup.GET("/reward_points_settings", lc.Fetch)
	itemGroup.GET("/reward_points_setting", lc.FetchByID)
	itemGroup.POST("/reward_points_setting", lc.Create)
	itemGroup.POST("/reward_points_settings", lc.CreateMany)
	itemGroup.PUT("/reward_points_setting", lc.Update)
	itemGroup.DELETE("/reward_points_setting", lc.Delete)
}
