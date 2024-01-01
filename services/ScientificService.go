package services

import (
	"errors"
	"fmt"
	"goblastapi-example/models"

	"github.com/mrkresnofatih/goblast"
)

type IScientificService interface {
	Exponent(request goblast.ContextfulReq[models.ScientificRequest]) (models.ScientificResponse, error)
	Combination(request goblast.ContextfulReq[models.ScientificRequest]) (models.ScientificResponse, error)
	Permutation(request goblast.ContextfulReq[models.ScientificRequest]) (models.ScientificResponse, error)
}

type ScientificService struct {
}

func (s *ScientificService) Exponent(request goblast.ContextfulReq[models.ScientificRequest]) (models.ScientificResponse, error) {
	goblast.LogInf(request.GetMetadata(), fmt.Sprintf("start exponent w. data: %f ^ %f", request.ReqData.VariableOne, request.ReqData.VariableTwo))
	if int(request.ReqData.VariableOne) == 0 {
		goblast.LogErr(request.GetMetadata(), "exponent base cannot be zero")
		return *new(models.ScientificResponse), errors.New("exponent base cannot be zero")
	}
	if int(request.ReqData.VariableTwo) == 0 {
		return models.ScientificResponse{
			Result: 1,
		}, nil
	}
	if int(request.ReqData.VariableTwo) < 0 {
		var negativeExpResult float32 = 1
		for i := 0; i < int(request.ReqData.VariableTwo); i++ {
			negativeExpResult = negativeExpResult / request.ReqData.VariableOne
		}
		return models.ScientificResponse{
			Result: negativeExpResult,
		}, nil
	}
	var positiveExpResult float32 = 1
	for i := 0; i < int(request.ReqData.VariableTwo); i++ {
		positiveExpResult = positiveExpResult * request.ReqData.VariableOne
	}
	return models.ScientificResponse{
		Result: positiveExpResult,
	}, nil
}

func (s *ScientificService) Combination(request goblast.ContextfulReq[models.ScientificRequest]) (models.ScientificResponse, error) {
	goblast.LogInf(request.GetMetadata(), fmt.Sprintf("start combination w. data: n=%d & r=%d", int(request.ReqData.VariableOne), int(request.ReqData.VariableTwo)))
	result := factorial(int(request.ReqData.VariableOne)) / (factorial(int(request.ReqData.VariableTwo)) * factorial(int(request.ReqData.VariableOne)-int(request.ReqData.VariableTwo)))
	return models.ScientificResponse{
		Result: float32(result),
	}, nil
}

func (s *ScientificService) Permutation(request goblast.ContextfulReq[models.ScientificRequest]) (models.ScientificResponse, error) {
	goblast.LogInf(request.GetMetadata(), fmt.Sprintf("start permutation w. data: n=%d & r=%d", int(request.ReqData.VariableOne), int(request.ReqData.VariableTwo)))
	result := factorial(int(request.ReqData.VariableOne)) / (factorial(int(request.ReqData.VariableOne) - int(request.ReqData.VariableTwo)))
	return models.ScientificResponse{
		Result: float32(result),
	}, nil
}

func factorial(num int) int {
	if num < 0 {
		return -1
	}
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
