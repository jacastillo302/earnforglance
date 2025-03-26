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

func ForumGroupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumGroupRepository(db, domain.CollectionForumGroup)
	lc := &controller.ForumGroupController{
		ForumGroupUsecase: usecase.NewForumGroupUsecase(ur, timeout),
		Env:               env,
	}

	itemGroup := group.Group("/api/v1/forums")
	itemGroup.GET("/forum_groups", lc.Fetch)
	itemGroup.GET("/forum_group", lc.FetchByID)
	itemGroup.POST("/forum_group", lc.Create)
	itemGroup.POST("/forum_groups", lc.CreateMany)
	itemGroup.PUT("/forum_group", lc.Update)
	itemGroup.DELETE("/forum_group", lc.Delete)
}
