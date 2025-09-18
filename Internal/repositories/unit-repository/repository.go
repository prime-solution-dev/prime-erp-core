package unitRepository

import (
	"prime-erp-core/internal/db"
	"prime-erp-core/internal/models"
)

func GetAllUnit() ([]models.Unit, error) {
	gormx, err := db.ConnectGORM(`prime_erp`)
	defer db.CloseGORM(gormx)
	if err != nil {
		return nil, err
	}

	units := []models.Unit{}

	result := gormx.Preload("UnitMethodItems").Preload("UnitMethodItems.UnitUomItems").Find(&units)
	if result.Error != nil {
		return nil, result.Error
	}

	return units, nil
}
