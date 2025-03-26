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

func PollVotingRecordRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPollVotingRecordRepository(db, domain.CollectionPollVotingRecord)
	lc := &controller.PollVotingRecordController{
		PollVotingRecordUsecase: usecase.NewPollVotingRecordUsecase(ur, timeout),
		Env:                     env,
	}

	itemGroup := group.Group("/api/v1/polls")

	itemGroup.GET("/poll_voting_records", lc.Fetch)
	itemGroup.GET("/poll_voting_record", lc.FetchByID)
	itemGroup.POST("/poll_voting_record", lc.Create)
	itemGroup.POST("/poll_voting_records", lc.CreateMany)
	itemGroup.PUT("/poll_voting_record", lc.Update)
	itemGroup.DELETE("/poll_voting_record", lc.Delete)
}
