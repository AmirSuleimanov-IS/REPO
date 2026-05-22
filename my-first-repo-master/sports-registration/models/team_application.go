package models

import "time"

type Application struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null;index:idx_user_draft,where:status='draft'"`
	Status      string    `gorm:"size:20;default:'draft';check:status IN ('draft','deleted','formed','completed','rejected')"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	FormedAt    *time.Time
	CompletedAt *time.Time
	ModeratorID *uint
	TeamName    string    `gorm:"size:100;not null"`
	TotalAmount float64   `gorm:"type:decimal(10,2);default:0.00"`

	Services []ApplicationService `gorm:"foreignKey:ApplicationID"`
}

func (Application) TableName() string {
	return "applications"
}
