package route

import (
	"time"

	controller "earnforglance/server/api/controller/vendors"
	"earnforglance/server/bootstrap"
	domain "earnforglance/server/domain/vendors"

	"earnforglance/server/mongo"
	repository "earnforglance/server/repository/vendors"
	usecase "earnforglance/server/usecase/vendors"

	"github.com/gin-gonic/gin"
)

func VendorNoteRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewVendorNoteRepository(db, domain.CollectionVendorNote)
	lc := &controller.VendorNoteController{
		VendorNoteUsecase: usecase.NewVendorNoteUsecase(ur, timeout),
		Env:               env,
	}

	group.GET("/vendor_notes", lc.Fetch)
	group.GET("/vendor_note", lc.FetchByID)
	group.POST("/vendor_note", lc.Create)
	group.POST("/vendor_notes", lc.CreateMany)
	group.PUT("/vendor_note", lc.Update)
	group.DELETE("/vendor_note", lc.Delete)
}
