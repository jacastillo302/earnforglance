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

	group.GET("/search_terms", lc.Fetch)
	group.GET("/search_term", lc.FetchByID)
	group.POST("/search_term", lc.Create)
	group.PUT("/search_term", lc.Update)
	group.DELETE("/search_term", lc.Delete)
}
