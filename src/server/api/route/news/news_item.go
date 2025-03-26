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

func NewsItemRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewNewsItemRepository(db, domain.CollectionNewsItem)
	lc := &controller.NewsItemController{
		NewsItemUsecase: usecase.NewNewsItemUsecase(ur, timeout),
		Env:             env,
	}
	itemGroup := group.Group("/api/v1/news")
	itemGroup.GET("/newsitems", lc.Fetch)
	itemGroup.GET("/newsitem", lc.FetchByID)
	itemGroup.POST("/newsitem", lc.Create)
	itemGroup.POST("/newsitems", lc.CreateMany)
	itemGroup.PUT("/newsitem", lc.Update)
	itemGroup.DELETE("/newsitem", lc.Delete)
}
