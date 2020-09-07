package request

// User scheme
type User struct {
	FirstName string `json:"firstName,omitempty" example:"Edos"`
	LastName  string `json:"lastName,omitempty" example:"Gaifut"`
	Email     string `json:"email,omitempty" example:"razor3538@mail.ru"`
}
