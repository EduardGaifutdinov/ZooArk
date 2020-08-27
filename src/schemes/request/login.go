package request

//LoginUserRequest user model for /login request route
type LoginUserRequest struct {
	Email    string `json:"email" example:"admin@mail.ru" binding:"required"`
	Password string `json:"password" example:"Password12!" binding:"required"`
} //@name LoginRequest
