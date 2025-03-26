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

func ForumPostRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumPostRepository(db, domain.CollectionForumPost)
	lc := &controller.ForumPostController{
		ForumPostUsecase: usecase.NewForumPostUsecase(ur, timeout),
		Env:              env,
	}

	itemGroup := group.Group("/api/v1/forums")
	itemGroup.GET("/forumposts", lc.Fetch)
	itemGroup.GET("/forumpost", lc.FetchByID)
	itemGroup.POST("/forumpost", lc.Create)
	itemGroup.POST("/forumposts", lc.CreateMany)
	itemGroup.PUT("/forumpost", lc.Update)
	itemGroup.DELETE("/forumpost", lc.Delete)
}
