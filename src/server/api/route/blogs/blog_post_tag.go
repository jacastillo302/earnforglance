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

func BlogPostTagRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBlogPostTagRepository(db, domain.CollectionBlogPostTag)
	lc := &controller.BlogPostTagController{
		BlogPostTagUsecase: usecase.NewBlogPostTagUsecase(ur, timeout),
		Env:                env,
	}

	group.GET("/blog_post_tags", lc.Fetch)
	group.GET("/blog_post_tag", lc.FetchByID)
	group.POST("/blog_post_tag", lc.Create)
	group.POST("/blog_post_tags", lc.CreateMany)
	group.PUT("/blog_post_tag", lc.Update)
	group.DELETE("blog_post_tag", lc.Delete)
}
