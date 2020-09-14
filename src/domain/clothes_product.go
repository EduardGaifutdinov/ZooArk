package domain

import uuid "github.com/satori/go.uuid"

// clothes struct
type Clothes struct {
	Base
	Name       string    `gorm:"type:varchar(100); not null" json:"name"`
	Count      int       `sql:"NullInt64: not null" json:"count"`
	Price      int       `sql:"NullInt64: not null" json:"price"`
	CategoryID uuid.UUID `json:"-"`
	Type       string    `gorm:"type:varchar(100); not null" json:"type"`
	Color      string    `gorm:"type:varchar(100); not null" json:"color"`
	Stock      string    `gorm:"type:varchar(10); not null" json:"stock"`
}
