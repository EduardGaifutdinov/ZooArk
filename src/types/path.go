package types

// PathID struct fo path binding
type PathID struct {
	ID string `url:"id" json:"id" binding:"required"`
} // @name IDResponse

// PathUser struct for path binding
type PathUser struct {
	Id string `uri:"id" json:"id" binding:"required"`
}
