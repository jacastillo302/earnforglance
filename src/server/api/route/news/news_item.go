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

	group.GET("/newsitems", lc.Fetch)
	group.GET("/newsitem", lc.FetchByID)
	group.POST("/newsitem", lc.Create)
	group.POST("/newsitems", lc.CreateMany)
	group.PUT("/newsitem", lc.Update)
	group.DELETE("/newsitem", lc.Delete)
}
