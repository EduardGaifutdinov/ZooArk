package types

type enum struct {
	SuperAdmin string
	User string
}

// UserRoleEnum enum
var UserRoleEnum = enum{
	SuperAdmin: "Super administrator",
	User: "User",
}
