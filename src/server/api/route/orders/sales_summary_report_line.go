package route

import (
	"time"

	controller "earnforglance/server/api/controller/orders"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/orders"

	repository "earnforglance/server/repository/orders"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/orders"

	"github.com/gin-gonic/gin"
)

func SalesSummaryReportLineRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewSalesSummaryReportLineRepository(db, domain.CollectionSalesSummaryReportLine)
	lc := &controller.SalesSummaryReportLineController{
		SalesSummaryReportLineUsecase: usecase.NewSalesSummaryReportLineUsecase(ur, timeout),
		Env:                           env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/sales_summary_report_lines", lc.Fetch)
	itemGroup.GET("/sales_summary_report_line", lc.FetchByID)
	itemGroup.POST("/sales_summary_report_line", lc.Create)
	itemGroup.POST("/sales_summary_report_lines", lc.CreateMany)
	itemGroup.PUT("/sales_summary_report_line", lc.Update)
	itemGroup.DELETE("/sales_summary_report_line", lc.Delete)
}
