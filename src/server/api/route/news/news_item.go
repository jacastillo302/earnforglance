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

func NewsItemRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsItemRepository(db, domain.CollectionNewsItem)
	lc := &controller.NewsItemController{
		NewsItemUsecase: usecase.NewNewsItemUsecase(ur, timeout),
		Env:             env,
	}
	itemGroup := group.Group("/api/v1/news")
	itemGroup.GET("/news_items", lc.Fetch)
	itemGroup.GET("/news_item", lc.FetchByID)
	itemGroup.POST("/news_item", lc.Create)
	itemGroup.POST("/news_items", lc.CreateMany)
	itemGroup.PUT("/news_item", lc.Update)
	itemGroup.DELETE("/news_item", lc.Delete)
}
