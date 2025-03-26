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

	itemGroup := group.Group("/api/v1/customers")
	itemGroup.GET("/reward_points_histories", lc.Fetch)
	itemGroup.GET("/reward_points_history", lc.FetchByID)
	itemGroup.POST("/reward_points_history", lc.Create)
	itemGroup.POST("/reward_points_histories", lc.CreateMany)
	itemGroup.PUT("/reward_points_history", lc.Update)
	itemGroup.DELETE("/reward_points_history", lc.Delete)
}
