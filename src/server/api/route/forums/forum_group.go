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

	group.GET("/forum_groups", lc.Fetch)
	group.GET("/forum_group", lc.FetchByID)
	group.POST("/forum_group", lc.Create)
	group.PUT("/forum_group", lc.Update)
	group.DELETE("/forum_group", lc.Delete)
}
