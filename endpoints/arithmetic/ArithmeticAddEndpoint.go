package arithmetic

import (
	"fmt"
	"goblastapi-example/models"
	"goblastapi-example/services"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mrkresnofatih/goblast"
)

type ArithmeticAddEndpoint struct {
	ArithmeticService services.IArithmeticService
}

func (a *ArithmeticAddEndpoint) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		addReq := new(goblast.ContextfulReq[models.ArithmeticRequest])
		err := c.Bind(addReq)
		if err != nil {
			goblast.LogError(uuid.NewString(), uuid.NewString(), "failed to parse to model")
			return goblast.NotOkResponse(c, "failed to parse model")
		}

		goblast.LogInf(addReq.GetMetadata(), fmt.Sprintf("data: varOne = %f && varTwo = %f", addReq.ReqData.VariableOne, addReq.ReqData.VariableTwo))
		result, err := a.ArithmeticService.Add(*addReq)
		if err != nil {
			goblast.LogErr(addReq.GetMetadata(), "something went wrong while adding")
			return goblast.NotOkResponse(c, "failed to add")
		}

		return goblast.OkResponse[models.ArithmeticResponse](c, result)
	}
}

func (a *ArithmeticAddEndpoint) GetPath() string {
	return "/add"
}

func (a *ArithmeticAddEndpoint) Register(group *echo.Group) {
	goblast.RegisterEndpointHelper(a, group)
}
