package models

type Planet struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"size:255" json:"name"`
	Population     string `gorm:"size:255" json:"population"`
	Diameter       string `gorm:"size:255" json:"diameter"`
	RotationPeriod string `gorm:"size:255" json:"rotation_period"`
	OrbitalPeriod  string `gorm:"size:255" json:"orbital_period"`
}

type APIPlanet struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"size:255" json:"name"`
	Population     string `gorm:"size:255" json:"population"`
	Diameter       string `gorm:"size:255" json:"diameter"`
	RotationPeriod string `gorm:"size:255" json:"rotation_period"`
	OrbitalPeriod  string `gorm:"size:255" json:"orbital_period"`
}

func (Planet) TableName() string {
	return "planet"
}
