package types

// PaginationQuery struct user for pagination binding
type PaginationQuery struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

// UserFilterQuery struct for filter and sort users in DB
type UserFilterQuery struct {
	Query  string `form:"q"`
	Role   string `form:"role"`
	Status string `form:"status"`
}
