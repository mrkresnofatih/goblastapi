package services

import (
	"fmt"
	"goblastapi-example/models"

	"github.com/mrkresnofatih/goblast"
)

type IArithmeticService interface {
	Add(request goblast.ContextfulReq[models.ArithmeticRequest]) (models.ArithmeticResponse, error)
	Subtract(request goblast.ContextfulReq[models.ScientificRequest]) (models.ArithmeticResponse, error)
}

type ArithmeticService struct {
}

func (a ArithmeticService) Add(request goblast.ContextfulReq[models.ArithmeticRequest]) (models.ArithmeticResponse, error) {
	goblast.LogInf(request.GetMetadata(), fmt.Sprintf("start execute add w. data: %f + %f", request.ReqData.VariableOne, request.ReqData.VariableTwo))
	result := request.ReqData.VariableOne + request.ReqData.VariableTwo
	return models.ArithmeticResponse{
		Result: result,
	}, nil
}

func (a ArithmeticService) Subtract(request goblast.ContextfulReq[models.ScientificRequest]) (models.ArithmeticResponse, error) {
	goblast.LogInf(request.GetMetadata(), fmt.Sprintf("start execute subtract w. data: %f - %f", request.ReqData.VariableOne, request.ReqData.VariableTwo))
	result := request.ReqData.VariableOne - request.ReqData.VariableTwo
	return models.ArithmeticResponse{
		Result: result,
	}, nil
}
