package route

import (
	"time"

	controller "earnforglance/server/api/controller/vendors"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/vendors"

	repository "earnforglance/server/repository/vendors"
	"earnforglance/server/service/data/mongo"
	usecase "earnforglance/server/usecase/vendors"

	"github.com/gin-gonic/gin"
)

func VendorNoteRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVendorNoteRepository(db, domain.CollectionVendorNote)
	lc := &controller.VendorNoteController{
		VendorNoteUsecase: usecase.NewVendorNoteUsecase(ur, timeout),
		Env:               env,
	}
	itemGroup := group.Group("/api/v1/vendors")
	itemGroup.GET("/vendor_notes", lc.Fetch)
	itemGroup.GET("/vendor_note", lc.FetchByID)
	itemGroup.POST("/vendor_note", lc.Create)
	itemGroup.POST("/vendor_notes", lc.CreateMany)
	itemGroup.PUT("/vendor_note", lc.Update)
	itemGroup.DELETE("/vendor_note", lc.Delete)
}
