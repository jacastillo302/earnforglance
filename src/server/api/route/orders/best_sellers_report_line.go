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

func BestSellersReportLineRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewBestSellersReportLineRepository(db, domain.CollectionBestSellersReportLine)
	lc := &controller.BestSellersReportLineController{
		BestSellersReportLineUsecase: usecase.NewBestSellersReportLineUsecase(ur, timeout),
		Env:                          env,
	}
	itemGroup := group.Group("/api/v1/orders")
	itemGroup.GET("/best_sellers_report_lines", lc.Fetch)
	itemGroup.GET("/best_sellers_report_line", lc.FetchByID)
	itemGroup.POST("/best_sellers_report_line", lc.Create)
	itemGroup.POST("/best_sellers_report_lines", lc.CreateMany)
	itemGroup.PUT("/best_sellers_report_line", lc.Update)
	itemGroup.DELETE("/best_sellers_report_line", lc.Delete)
}
