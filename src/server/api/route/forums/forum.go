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

func ForumRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewForumRepository(db, domain.CollectionForum)
	lc := &controller.ForumController{
		ForumUsecase: usecase.NewForumUsecase(ur, timeout),
		Env:          env,
	}

	group.GET("/forums", lc.Fetch)
	group.GET("/forum", lc.FetchByID)
	group.POST("/forum", lc.Create)
	group.PUT("/forum", lc.Update)
	group.DELETE("/forum", lc.Delete)
}
