package services

import (
	"encoding/json"
	"strconv"
	"template_api/database"
	"template_api/models"
	"time"
)

func GetPeopleFromDB(offset int) ([]models.Person, int64) {
	var people []models.Person
	var count int64

	database.Database.Offset(offset).Limit(database.PageLimit).Find(&people)

	database.Database.Model(models.Person{}).Count(&count)

	return people, count
}

func GetPersonById(personId int) models.Person {
	var person models.Person

	personRedis, err := database.Redis.Get(database.RedisCtx, "person_"+strconv.Itoa(personId)).Result()
	if err != nil {
		person = GetPersonByIdFromDB(personId)
		data, err := json.Marshal(person)
		err = database.Redis.Set(database.RedisCtx, "person_"+strconv.Itoa(personId), data, 5*time.Minute).Err()
		if err != nil {
			panic(err)
		}
	} else {
		err = json.Unmarshal([]byte(personRedis), &person)
		if err != nil {
			panic(err)
		}
	}
	return person
}

func GetPersonByIdFromDB(personId int) models.Person {
	var person models.Person

	database.Database.First(&person, personId)

	return person
}
