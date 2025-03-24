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

func ForumPostRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumPostRepository(db, domain.CollectionForumPost)
	lc := &controller.ForumPostController{
		ForumPostUsecase: usecase.NewForumPostUsecase(ur, timeout),
		Env:              env,
	}

	group.GET("/forumposts", lc.Fetch)
	group.GET("/forumpost", lc.FetchByID)
	group.POST("/forumpost", lc.Create)
	group.POST("/forumposts", lc.CreateMany)
	group.PUT("/forumpost", lc.Update)
	group.DELETE("/forumpost", lc.Delete)
}
