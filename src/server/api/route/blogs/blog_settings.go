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

func BlogSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBlogSettingsRepository(db, domain.CollectionBlogSettings)
	lc := &controller.BlogSettingsController{
		BlogSettingsUsecase: usecase.NewBlogSettingsUsecase(ur, timeout),
		Env:                 env,
	}

	itemGroup := group.Group("/api/v1/blogs")

	itemGroup.GET("/blog_settings", lc.Fetch)
	itemGroup.GET("/blog_setting", lc.FetchByID)
	itemGroup.POST("/blog_setting", lc.Create)
	itemGroup.POST("/blog_settings", lc.CreateMany)
	itemGroup.PUT("/blog_setting", lc.Update)
	itemGroup.DELETE("blog_setting", lc.Delete)
}
