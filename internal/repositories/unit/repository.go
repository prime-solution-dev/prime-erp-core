package unitRepository

import (
	"prime-erp-core/internal/db"
	"prime-erp-core/internal/models"
)

func GetAllUnit() ([]models.Unit, error) {
	gormx, err := db.ConnectGORM("prime_erp")
	if err != nil {
		return nil, err
	}
	defer db.CloseGORM(gormx)

	units := []models.Unit{}

	result := gormx.Preload("UnitMethodItems").Preload("UnitMethodItems.UnitUomItems").Find(&units)
	if result.Error != nil {
		return nil, result.Error
	}

	return units, nil
}
