package request

import (
	//"fmt"
	. "github.com/space-tech-dev/sms-service-notification/app/http"
	. "github.com/space-tech-dev/sms-service-notification/config"
	//"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/gin-gonic/gin"
)

type ParaPushReservationWaitToReservedPlace struct {
	ReservationUUID string `form:"reservationUUID" json:"reservationUUID" xml:"reservationUUID"  binding:"required"`
}

func ValidatePushReservationWaitToReservedPlace(context *gin.Context) {
	var form ParaPushReservationWaitToReservedPlace

	//fmt.Println("validate notification push reserve normal to wait list")
	if err := context.ShouldBind(&form); err != nil {
		//println(err.Error())
		if err := context.ShouldBindJSON(&form); err != nil {
			apiResponse := &APIResponse{}
			apiResponse.SetCode(
				API_ERR_CODE_REQUEST_PARAM_ERROR,
				API_RETURN_CODE_ERROR,
				"", "").SetData(object.HashMap{
				"message": err.Error(),
			}).ThrowJSONResponse(context)
		}
	}

	context.Set("params", form)

	context.Next()
}
