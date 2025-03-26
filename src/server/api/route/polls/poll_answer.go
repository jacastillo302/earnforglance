package route

import (
	"time"

	controller "earnforglance/server/api/controller/polls"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/polls"

	repository "earnforglance/server/repository/polls"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/polls"

	"github.com/gin-gonic/gin"
)

func PollAnswerRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPollAnswerRepository(db, domain.CollectionPollAnswer)
	lc := &controller.PollAnswerController{
		PollAnswerUsecase: usecase.NewPollAnswerUsecase(ur, timeout),
		Env:               env,
	}
	itemGroup := group.Group("/api/v1/polls")
	itemGroup.GET("/poll_answers", lc.Fetch)
	itemGroup.GET("/poll_answer", lc.FetchByID)
	itemGroup.POST("/poll_answer", lc.Create)
	itemGroup.POST("/poll_answers", lc.CreateMany)
	itemGroup.PUT("/poll_answer", lc.Update)
	itemGroup.DELETE("/poll_answer", lc.Delete)
}
