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

func BlogPostTagRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBlogPostTagRepository(db, domain.CollectionBlogPostTag)
	lc := &controller.BlogPostTagController{
		BlogPostTagUsecase: usecase.NewBlogPostTagUsecase(ur, timeout),
		Env:                env,
	}

	itemGroup := group.Group("/api/v1/blogs")

	itemGroup.GET("/blog_post_tags", lc.Fetch)
	itemGroup.GET("/blog_post_tag", lc.FetchByID)
	itemGroup.POST("/blog_post_tag", lc.Create)
	itemGroup.POST("/blog_post_tags", lc.CreateMany)
	itemGroup.PUT("/blog_post_tag", lc.Update)
	itemGroup.DELETE("/blog_post_tag", lc.Delete)
}
