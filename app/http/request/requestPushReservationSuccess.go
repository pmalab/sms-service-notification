package request

import (
	//"fmt"
	. "github.com/space-tech-dev/sms-service-notification/app/http"
	. "github.com/space-tech-dev/sms-service-notification/config"
	//"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/gin-gonic/gin"
)

type ParaPushReservationSuccess struct {
	ReservationUUID string `form:"reservationUUID" json:"reservationUUID" xml:"reservationUUID"  binding:"required"`
}

func ValidatePushReservationSuccess(context *gin.Context) {
	var form ParaPushReservationSuccess

	//fmt.Println("validate notification push reserve success")
	if err := context.ShouldBind(&form); err != nil {
		//println(err.Error())
		if err := context.ShouldBindJSON(&form); err != nil {
			//println("Error occurs", err.Error())
			//js, _ := json.Marshal(gin.H{"error": err.Error()})
			//context.Writer.Header().Set("Content-Type", "application/json")
			//context.Writer.Write(js)
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
