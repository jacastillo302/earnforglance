package route

import (
	"time"

	controller "earnforglance/server/api/controller/news"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/news"

	repository "earnforglance/server/repository/news"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/news"

	"github.com/gin-gonic/gin"
)

func NewsCommentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsCommentRepository(db, domain.CollectionNewsComment)
	lc := &controller.NewsCommentController{
		NewsCommentUsecase: usecase.NewNewsCommentUsecase(ur, timeout),
		Env:                env,
	}
	itemGroup := group.Group("/api/v1/news")
	itemGroup.GET("/news_comments", lc.Fetch)
	itemGroup.GET("/news_comment", lc.FetchByID)
	itemGroup.POST("/news_comment", lc.Create)
	itemGroup.POST("/news_comments", lc.CreateMany)
	itemGroup.PUT("/news_comment", lc.Update)
	itemGroup.DELETE("/news_comment", lc.Delete)
}
