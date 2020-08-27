package types

type statusEnem struct {
	Active  string
	Invited string
	Deleted string
}

// StatusTypesEnum enum
var StatusTypesEnum = statusEnem{
	Active:  "active",
	Invited: "invited",
	Deleted: "deleted",
}
