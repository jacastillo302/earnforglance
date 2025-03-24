package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/orders"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func SalesSummaryReportLineRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSalesSummaryReportLineRepository(db, domain.CollectionSalesSummaryReportLine)
	lc := &controller.SalesSummaryReportLineController{
		SalesSummaryReportLineUsecase: usecase.NewSalesSummaryReportLineUsecase(ur, timeout),
		Env:                           env,
	}

	group.GET("/sales_summary_report_lines", lc.Fetch)
	group.GET("/sales_summary_report_line", lc.FetchByID)
	group.POST("/sales_summary_report_line", lc.Create)
	group.POST("/sales_summary_report_lines", lc.CreateMany)
	group.PUT("/sales_summary_report_line", lc.Update)
	group.DELETE("/sales_summary_report_line", lc.Delete)
}
