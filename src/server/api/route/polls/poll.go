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

	group.GET("/polls", lc.Fetch)
	group.GET("/poll", lc.FetchByID)
	group.POST("/poll", lc.Create)
	group.POST("/polls", lc.CreateMany)
	group.PUT("/poll", lc.Update)
	group.DELETE("/poll", lc.Delete)
}
