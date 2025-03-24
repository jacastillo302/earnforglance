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

func BlogSettingsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBlogSettingsRepository(db, domain.CollectionBlogSettings)
	lc := &controller.BlogSettingsController{
		BlogSettingsUsecase: usecase.NewBlogSettingsUsecase(ur, timeout),
		Env:                 env,
	}

	group.GET("/blog_settings", lc.Fetch)
	group.GET("/blog_setting", lc.FetchByID)
	group.POST("/blog_setting", lc.Create)
	group.POST("/blog_settings", lc.CreateMany)
	group.PUT("/blog_setting", lc.Update)
	group.DELETE("blog_setting", lc.Delete)
}
