package route

import (
	"time"

	controller "earnforglance/server/api/controller/news"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/news"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/news"
	usecase "earnforglance/server/usecase/news"

	"github.com/gin-gonic/gin"
)

func NewsCommentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsCommentRepository(db, domain.CollectionNewsComment)
	lc := &controller.NewsCommentController{
		NewsCommentUsecase: usecase.NewNewsCommentUsecase(ur, timeout),
		Env:                env,
	}

	group.GET("/news_comments", lc.Fetch)
	group.GET("/news_comment", lc.FetchByID)
	group.POST("/news_comment", lc.Create)
	group.PUT("/news_comment", lc.Update)
	group.DELETE("/news_comment", lc.Delete)
}
