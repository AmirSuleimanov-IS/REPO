package models

type Service struct {
	ID              uint    `gorm:"primaryKey;autoIncrement"`
	Name            string  `gorm:"size:150;not null"`
	Description     string  `gorm:"type:text"`
	Status          string  `gorm:"size:20;default:'active';check:status IN ('active','deleted')"`
	ImageKey        string  `gorm:"size:255"`
	VideoKey        string  `gorm:"size:255"`
	EventDate       string  `gorm:"size:20;not null"`
	Location        string  `gorm:"size:100;not null"`
	BasePrice       float64 `gorm:"type:decimal(10,2);not null"`
	Category        string  `gorm:"size:50;not null"`
	MaxParticipants int     `gorm:"not null"`
}

func (Service) TableName() string {
	return "services"
}
