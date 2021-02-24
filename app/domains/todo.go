package domains

import "time"

type (
	Todo struct {
		ID        int    `gorm:"primaryKey"`
		Name      string `gorm:"size:255;not null"`
		Status    string `gorm:"type:enum('yet','doing','done');default: 'yet';not null"`
		CreatedAt *time.Time
		UpdatedAt *time.Time
		DeletedAt *time.Time
	}
	Todos []Todo
)
