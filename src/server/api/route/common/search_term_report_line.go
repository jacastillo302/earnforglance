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

	group.GET("/search_term_report_lines", lc.Fetch)
	group.GET("/search_term_report_line", lc.FetchByID)
	group.POST("/search_term_report_line", lc.Create)
	group.PUT("/search_term_report_line", lc.Update)
	group.DELETE("/search_term_report_line", lc.Delete)
}
