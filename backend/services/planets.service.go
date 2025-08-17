package services

import (
	"template_api/database"
	"template_api/models"
)

func GetPlanetsFromDB(offset int) ([]models.Planet, int64) {
	var planets []models.Planet
	var count int64

	database.Database.Offset(offset).Limit(database.PageLimit).Find(&planets)

	database.Database.Model(models.Planet{}).Count(&count)

	return planets, count
}

func GetPlanetByIdFromDB(planetId int) models.Planet {
	var planet models.Planet

	database.Database.First(&planet, planetId)

	return planet
}
