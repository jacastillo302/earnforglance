package route

import (
	"time"

	controller "earnforglance/server/api/controller/blogs"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/blogs"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/blogs"
	usecase "earnforglance/server/usecase/blogs"

	"github.com/gin-gonic/gin"
)

func BlogCommentRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBlogCommentRepository(db, domain.CollectionBlogComment)
	lc := &controller.BlogCommentController{
		BlogCommentUsecase: usecase.NewBlogCommentUsecase(ur, timeout),
		Env:                env,
	}

	itemGroup := group.Group("/api/v1/blogs")

	itemGroup.GET("/blog_comments", lc.Fetch)
	itemGroup.GET("/blog_comment", lc.FetchByID)
	itemGroup.POST("/blog_comment", lc.Create)
	itemGroup.POST("/blog_comments", lc.CreateMany)
	itemGroup.PUT("/blog_comment", lc.Update)
	itemGroup.DELETE("/blog_comment", lc.Delete)
}
