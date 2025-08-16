package models

type Person struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"size:255" json:"name"`
	Gender    string `gorm:"size:255" json:"gender"`
	Height    uint   `json:"height"`
	Mass      uint   `json:"mass"`
	HairColor string `gorm:"size:255" json:"hair_color"`
	PlanetID  int    `json:"planet_id"`
	Planet    Planet
	URL       string `gorm:"size:255" json:"url"`
}

type APIPerson struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"size:255" json:"name"`
	Gender    string `gorm:"size:255" json:"gender"`
	Height    uint   `json:"height"`
	Mass      uint   `json:"mass"`
	HairColor string `gorm:"size:255" json:"hair_color"`
	Homeworld string `gorm:"size:255" json:"homeworld"`
	URL       string `gorm:"size:255" json:"url"`
}

func (Person) TableName() string {
	return "people"
}
