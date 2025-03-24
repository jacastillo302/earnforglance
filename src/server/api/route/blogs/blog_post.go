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

func BlogPostRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBlogPostRepository(db, domain.CollectionBlogPost)
	lc := &controller.BlogPostController{
		BlogPostUsecase: usecase.NewBlogPostUsecase(ur, timeout),
		Env:             env,
	}

	group.GET("/blog_posts", lc.Fetch)
	group.GET("/blog_post", lc.FetchByID)
	group.POST("/blog_post", lc.Create)
	group.POST("/blog_posts", lc.CreateMany)
	group.PUT("/blog_post", lc.Update)
	group.DELETE("/blog_post", lc.Delete)
}
