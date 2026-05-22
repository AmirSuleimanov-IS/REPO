package models

type ApplicationService struct {
	ApplicationID uint    `gorm:"primaryKey;index:idx_app_svc,unique:composite;constraint:OnDelete:RESTRICT"`
	ServiceID     uint    `gorm:"primaryKey;index:idx_app_svc,unique:composite;constraint:OnDelete:RESTRICT"`
	Quantity      int     `gorm:"not null;default:1"`
	RoleInEvent   string  `gorm:"size:50;default:'participant'"`
	FinalPrice    float64 `gorm:"type:decimal(10,2)"`

	Service Service `gorm:"foreignKey:ServiceID"`
}

func (ApplicationService) TableName() string {
	return "application_services"
}
