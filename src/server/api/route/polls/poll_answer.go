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

func PollAnswerRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPollAnswerRepository(db, domain.CollectionPollAnswer)
	lc := &controller.PollAnswerController{
		PollAnswerUsecase: usecase.NewPollAnswerUsecase(ur, timeout),
		Env:               env,
	}

	group.GET("/poll_answers", lc.Fetch)
	group.GET("/poll_answer", lc.FetchByID)
	group.POST("/poll_answer", lc.Create)
	group.POST("/poll_answers", lc.CreateMany)
	group.PUT("/poll_answer", lc.Update)
	group.DELETE("/poll_answer", lc.Delete)
}
