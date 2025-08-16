package database

import (
	"context"
	"template_api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"

	"log"
	"os"
	"time"
)

var Database *gorm.DB
var Redis *redis.Client
var RedisCtx = context.Background()

func Connect() error {
	var err error

	var DatabaseUri string = os.Getenv("APP_DB_USER") + ":" + os.Getenv("APP_DB_PASS") + "@" +
		"tcp(" + os.Getenv("APP_DB_HOST") + ":3306)/" + os.Getenv("APP_DB_NAME") +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	logLevel := logger.Silent
	if os.Getenv("APP_DEBUG") == "1" {
		logLevel = logger.Error
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	Database, err = gorm.Open(mysql.Open(DatabaseUri), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("APP_REDIS_HOST") + ":6379",
		Password: os.Getenv("APP_REDIS_PASS"),
		DB:       0, // use default DB
	})

	sqlDB, err := Database.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	autoMigrateAndSeed()

	return nil
}

func autoMigrateAndSeed() {
	var err error
	if err = Database.AutoMigrate(&models.Person{}, &models.Planet{}); err != nil {
		fiberlog.Error("Unable to AutoMigrate models")
	}

	// Seed planets if table is empty
	var planetCount int64
	if Database.Model(&models.Planet{}).Count(&planetCount); planetCount == 0 {
		planets := []models.Planet{
			{
				ID:             1,
				Name:           "Tatooine",
				Population:     "200000",
				Diameter:       "10465",
				RotationPeriod: "23",
				OrbitalPeriod:  "304",
			},
			{
				ID:             2,
				Name:           "Alderaan",
				Population:     "2000000000",
				Diameter:       "12500",
				RotationPeriod: "24",
				OrbitalPeriod:  "364",
			},
			{
				ID:             3,
				Name:           "Kashyyyk",
				Population:     "45000000",
				Diameter:       "12765",
				RotationPeriod: "26",
				OrbitalPeriod:  "381",
			},
			{
				ID:             4,
				Name:           "Coruscant",
				Population:     "1000000000000",
				Diameter:       "12240",
				RotationPeriod: "24",
				OrbitalPeriod:  "368",
			},
			{
				ID:             5,
				Name:           "Naboo",
				Population:     "4500000000",
				Diameter:       "12120",
				RotationPeriod: "26",
				OrbitalPeriod:  "312",
			},
		}
		if err := Database.Create(&planets).Error; err != nil {
			fiberlog.Error("Unable to seed planets")
		}
	}

	// Seed people if table is empty
	var peopleCount int64
	if Database.Model(&models.Person{}).Count(&peopleCount); peopleCount == 0 {
		people := []models.Person{
			{
				ID:        1,
				Name:      "Luke Skywalker",
				Gender:    "male",
				Height:    172,
				Mass:      77,
				HairColor: "blond",
				PlanetID:  1,
				URL:       "https://swapi.dev/api/people/1/",
			},
			{
				ID:        2,
				Name:      "Leia Organa",
				Gender:    "female",
				Height:    150,
				Mass:      49,
				HairColor: "brown",
				PlanetID:  2,
				URL:       "https://swapi.dev/api/people/5/",
			},
			{
				ID:        3,
				Name:      "Han Solo",
				Gender:    "male",
				Height:    180,
				Mass:      80,
				HairColor: "brown",
				PlanetID:  1,
				URL:       "https://swapi.dev/api/people/14/",
			},
			{
				ID:        4,
				Name:      "Chewbacca",
				Gender:    "male",
				Height:    228,
				Mass:      112,
				HairColor: "brown",
				PlanetID:  3,
				URL:       "https://swapi.dev/api/people/13/",
			},
			{
				ID:        5,
				Name:      "Anakin Skywalker",
				Gender:    "male",
				Height:    188,
				Mass:      84,
				HairColor: "blond",
				PlanetID:  1,
				URL:       "https://swapi.dev/api/people/11/",
			},
			{
				ID:        6,
				Name:      "Padmé Amidala",
				Gender:    "female",
				Height:    165,
				Mass:      45,
				HairColor: "brown",
				PlanetID:  5,
				URL:       "https://swapi.dev/api/people/35/",
			},
			{
				ID:        7,
				Name:      "Mace Windu",
				Gender:    "male",
				Height:    188,
				Mass:      84,
				HairColor: "none",
				PlanetID:  4,
				URL:       "https://swapi.dev/api/people/51/",
			},
		}
		if err := Database.Create(&people).Error; err != nil {
			fiberlog.Error("Unable to seed people")
		}
	}
}
