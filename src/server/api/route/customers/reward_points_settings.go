package route

import (
	"time"

	controller "earnforglance/server/api/controller/customers"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/customers"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/customers"
	usecase "earnforglance/server/usecase/customers"

	"github.com/gin-gonic/gin"
)

func RewardPointsSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewRewardPointsSettingsRepository(db, domain.CollectionRewardPointsSettings)
	lc := &controller.RewardPointsSettingsController{
		RewardPointsSettingsUsecase: usecase.NewRewardPointsSettingsUsecase(ur, timeout),
		Env:                         env,
	}

	group.GET("/reward_points_settingss", lc.Fetch)
	group.GET("/reward_points_settings", lc.FetchByID)
	group.POST("/reward_points_settings", lc.Create)
	group.PUT("/reward_points_settings", lc.Update)
	group.DELETE("/reward_points_settings", lc.Delete)
}
