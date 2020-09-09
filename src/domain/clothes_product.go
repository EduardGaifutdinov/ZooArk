package domain

import uuid "github.com/satori/go.uuid"

// clothes struct
type Clothes struct {
	ProductID uuid.UUID `json:"-"`
	Type string `gorm:"type:varchar(100); not null" json:"type"`
	Color string `gorm:"type:varchar(100); not null" json:"color"`
	Stock string `gorm:"type:varchar(10); not null" json:"stock"`
}
