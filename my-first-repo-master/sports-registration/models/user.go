package models

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"size:50;uniqueIndex;not null"`
	Email        string `gorm:"size:100;uniqueIndex;not null"`
	PasswordHash string `gorm:"size:255;not null"`
	Role         string `gorm:"size:20;default:'athlete'"`
}

func (User) TableName() string {
	return "users"
}
