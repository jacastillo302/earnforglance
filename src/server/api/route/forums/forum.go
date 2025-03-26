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
	itemGroup := group.Group("/api/v1/forums")
	itemGroup.GET("/forums", lc.Fetch)
	itemGroup.GET("/forum", lc.FetchByID)
	itemGroup.POST("/forum", lc.Create)
	itemGroup.POST("/forums", lc.CreateMany)
	itemGroup.PUT("/forum", lc.Update)
	itemGroup.DELETE("/forum", lc.Delete)
}
