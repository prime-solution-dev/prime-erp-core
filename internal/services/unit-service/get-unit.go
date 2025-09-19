package unitService

import (
	"prime-erp-core/internal/models"
	unitRepository "prime-erp-core/internal/repositories/unit"

	"github.com/gin-gonic/gin"
)

func MapUnitUomToResponse(uom models.UnitUom) models.GetUnitUomResponse {
	return models.GetUnitUomResponse{
		ID:      uom.ID,
		UomCode: uom.UomCode,
		UomName: uom.UomName,
	}
}

func MapUnitMethodToResponse(modelMethod models.UnitMethod) models.GetUnitMethodResponse {
	uomResponses := make([]models.GetUnitUomResponse, len(modelMethod.UnitUomItems))
	for i, uom := range modelMethod.UnitUomItems {
		uomResponses[i] = MapUnitUomToResponse(uom)
	}

	return models.GetUnitMethodResponse{
		ID:           modelMethod.ID,
		MethodCode:   modelMethod.MethodCode,
		MethodName:   modelMethod.MethodName,
		UnitUomItems: uomResponses,
	}
}

func MapUnitToResponse(unit models.Unit) models.GetAllUnitResponse {
	methodResponses := make([]models.GetUnitMethodResponse, len(unit.UnitMethodItems))
	for i, method := range unit.UnitMethodItems {
		methodResponses[i] = MapUnitMethodToResponse(method)
	}

	return models.GetAllUnitResponse{
		ID:              unit.ID.String(),
		UnitCode:        unit.UnitCode,
		UnitName:        unit.UnitName,
		UnitMethodItems: methodResponses,
	}
}

func GetAllUnit(ctx *gin.Context, jsonPayload string) (interface{}, error) {
	units, err := unitRepository.GetAllUnit()
	if err != nil {
		return nil, err
	}

	unitResponses := make([]models.GetAllUnitResponse, len(units))
	for i, unit := range units {
		unitResponses[i] = MapUnitToResponse(unit)
	}

	return unitResponses, nil
}
