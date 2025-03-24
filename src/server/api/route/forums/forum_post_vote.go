package route

import (
	"time"

	controller "earnforglance/server/api/controller/forums"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/forums"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/forums"
	usecase "earnforglance/server/usecase/forums"

	"github.com/gin-gonic/gin"
)

func ForumPostVoteRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumPostVoteRepository(db, domain.CollectionForumPostVote)
	lc := &controller.ForumPostVoteController{
		ForumPostVoteUsecase: usecase.NewForumPostVoteUsecase(ur, timeout),
		Env:                  env,
	}

	group.GET("/forum_post_votes", lc.Fetch)
	group.GET("/forum_post_vote", lc.FetchByID)
	group.POST("/forum_post_vote", lc.Create)
	group.POST("/forum_post_votes", lc.CreateMany)
	group.PUT("/forum_post_vote", lc.Update)
	group.DELETE("/forum_post_vote", lc.Delete)
}
