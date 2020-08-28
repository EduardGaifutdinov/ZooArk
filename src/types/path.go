package types
 // PathID struct fo path binding
 type PathID struct {
 	ID string `url:"id" json:"id" binding:"required"`
 } // @name IDResponse