package services

import (
	"template_api/database"
	"template_api/models"
)

func GetPeopleFromDB(offset int) ([]models.Person, int64) {
	var people []models.Person
	var count int64

	database.Database.Offset(offset).Limit(database.PageLimit).Find(&people)

	database.Database.Model(models.Person{}).Count(&count)

	return people, count
}

func GetPersonByIdFromDB(personId int) models.Person {
	var person models.Person

	database.Database.First(&person, personId)

	return person
}
