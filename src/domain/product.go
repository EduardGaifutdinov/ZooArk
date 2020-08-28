package domain

// Product model
type Product struct {
	Base
	Name string `gorm:"type:varchar(100); not null" json:"name"`
	Count int `sql:"NullInt64: not null" json:"count"`
	Price int `sql:"NullInt64: not null" json:"price"`
}