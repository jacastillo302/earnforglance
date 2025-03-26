package route

import (
	"time"

	controller "earnforglance/server/api/controller/blogs"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/blogs"

	repository "earnforglance/server/repository/blogs"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/blogs"

	"github.com/gin-gonic/gin"
)

func BlogPostRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBlogPostRepository(db, domain.CollectionBlogPost)
	lc := &controller.BlogPostController{
		BlogPostUsecase: usecase.NewBlogPostUsecase(ur, timeout),
		Env:             env,
	}

	itemGroup := group.Group("/api/v1/blogs")

	itemGroup.GET("/blog_posts", lc.Fetch)
	itemGroup.GET("/blog_post", lc.FetchByID)
	itemGroup.POST("/blog_post", lc.Create)
	itemGroup.POST("/blog_posts", lc.CreateMany)
	itemGroup.PUT("/blog_post", lc.Update)
	itemGroup.DELETE("/blog_post", lc.Delete)
}
