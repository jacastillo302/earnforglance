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

func RewardPointsHistoryRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewRewardPointsHistoryRepository(db, domain.CollectionRewardPointsHistory)
	lc := &controller.RewardPointsHistoryController{
		RewardPointsHistoryUsecase: usecase.NewRewardPointsHistoryUsecase(ur, timeout),
		Env:                        env,
	}

	group.GET("/reward_points_histories", lc.Fetch)
	group.GET("/reward_points_history", lc.FetchByID)
	group.POST("/reward_points_history", lc.Create)
	group.PUT("/reward_points_history", lc.Update)
	group.DELETE("/reward_points_history", lc.Delete)
}
