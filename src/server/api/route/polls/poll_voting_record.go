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

func PollVotingRecordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPollVotingRecordRepository(db, domain.CollectionPollVotingRecord)
	lc := &controller.PollVotingRecordController{
		PollVotingRecordUsecase: usecase.NewPollVotingRecordUsecase(ur, timeout),
		Env:                     env,
	}

	group.GET("/poll_voting_records", lc.Fetch)
	group.GET("/poll_voting_record", lc.FetchByID)
	group.POST("/poll_voting_record", lc.Create)
	group.POST("/poll_voting_records", lc.CreateMany)
	group.PUT("/poll_voting_record", lc.Update)
	group.DELETE("/poll_voting_record", lc.Delete)
}
