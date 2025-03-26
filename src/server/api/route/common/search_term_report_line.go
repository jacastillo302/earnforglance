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

func SearchTermReportLineRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSearchTermReportLineRepository(db, domain.CollectionSearchTermReportLine)
	lc := &controller.SearchTermReportLineController{
		SearchTermReportLineUsecase: usecase.NewSearchTermReportLineUsecase(ur, timeout),
		Env:                         env,
	}

	itemGroup := group.Group("/api/v1/common")
	itemGroup.GET("/search_term_report_lines", lc.Fetch)
	itemGroup.GET("/search_term_report_line", lc.FetchByID)
	itemGroup.POST("/search_term_report_line", lc.Create)
	itemGroup.POST("/search_term_report_lines", lc.CreateMany)
	itemGroup.PUT("/search_term_report_line", lc.Update)
	itemGroup.DELETE("/search_term_report_line", lc.Delete)
}
