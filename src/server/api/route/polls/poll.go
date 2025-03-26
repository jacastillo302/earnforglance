package route

import (
	"time"

	controller "earnforglance/server/api/controller/polls"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/polls"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/polls"
	usecase "earnforglance/server/usecase/polls"

	"github.com/gin-gonic/gin"
)

func PollRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPollRepository(db, domain.CollectionPoll)
	lc := &controller.PollController{
		PollUsecase: usecase.NewPollUsecase(ur, timeout),
		Env:         env,
	}
	itemGroup := group.Group("/api/v1/polls")
	itemGroup.GET("/polls", lc.Fetch)
	itemGroup.GET("/poll", lc.FetchByID)
	itemGroup.POST("/poll", lc.Create)
	itemGroup.POST("/polls", lc.CreateMany)
	itemGroup.PUT("/poll", lc.Update)
	itemGroup.DELETE("/poll", lc.Delete)
}
