package models

type ScientificRequest struct {
	VariableOne float32 `json:"variableOne" validate:"required"`
	VariableTwo float32 `json:"variableTwo" validate:"required"`
}
