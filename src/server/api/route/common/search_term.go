package route

import (
	"time"

	controller "earnforglance/server/api/controller/common"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/common"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/common"
	usecase "earnforglance/server/usecase/common"

	"github.com/gin-gonic/gin"
)

func SearchTermRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSearchTermRepository(db, domain.CollectionSearchTerm)
	lc := &controller.SearchTermController{
		SearchTermUsecase: usecase.NewSearchTermUsecase(ur, timeout),
		Env:               env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/search_terms", lc.Fetch)
	itemGroup.GET("/search_term", lc.FetchByID)
	itemGroup.POST("/search_term", lc.Create)
	itemGroup.POST("/search_terms", lc.CreateMany)
	itemGroup.PUT("/search_term", lc.Update)
	itemGroup.DELETE("/search_term", lc.Delete)
}
