package services

import (
	"encoding/json"
	"strconv"
	"template_api/database"
	"template_api/models"
	"time"
)

func GetPlanetsFromDB(offset int) ([]models.Planet, int64) {
	var planets []models.Planet
	var count int64

	database.Database.Offset(offset).Limit(database.PageLimit).Find(&planets)

	database.Database.Model(models.Planet{}).Count(&count)

	return planets, count
}

func GetPlanetById(planetId int) models.Planet {
	var planet models.Planet

	planetRedis, err := database.Redis.Get(database.RedisCtx, "planet_"+strconv.Itoa(planetId)).Result()
	if err != nil {
		planet = GetPlanetByIdFromDB(planetId)
		data, err := json.Marshal(planet)
		err = database.Redis.Set(database.RedisCtx, "planet_"+strconv.Itoa(planetId), data, 5*time.Minute).Err()
		if err != nil {
			panic(err)
		}
	} else {
		err = json.Unmarshal([]byte(planetRedis), &planet)
		if err != nil {
			panic(err)
		}
	}
	return planet
}

func GetPlanetByIdFromDB(planetId int) models.Planet {
	var planet models.Planet

	database.Database.First(&planet, planetId)

	return planet
}
