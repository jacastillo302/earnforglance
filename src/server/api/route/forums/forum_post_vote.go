package route

import (
	"time"

	controller "earnforglance/server/api/controller/forums"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/forums"

	repository "earnforglance/server/repository/forums"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/forums"

	"github.com/gin-gonic/gin"
)

func ForumPostVoteRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumPostVoteRepository(db, domain.CollectionForumPostVote)
	lc := &controller.ForumPostVoteController{
		ForumPostVoteUsecase: usecase.NewForumPostVoteUsecase(ur, timeout),
		Env:                  env,
	}

	itemGroup := group.Group("/api/v1/forums")
	itemGroup.GET("/forum_post_votes", lc.Fetch)
	itemGroup.GET("/forum_post_vote", lc.FetchByID)
	itemGroup.POST("/forum_post_vote", lc.Create)
	itemGroup.POST("/forum_post_votes", lc.CreateMany)
	itemGroup.PUT("/forum_post_vote", lc.Update)
	itemGroup.DELETE("/forum_post_vote", lc.Delete)
}
