package request

// UserUpdate scheme
type UserUpdate struct {
	FirstName string `json:"firstName,omitempty" example:"Eduard"`
	LastName  string `json:"lastName,omitempty" example:"Gaifutdinov"`
	Email     string `json:"email,omitempty" example:"razzzor3538@mail.ru"`
}
