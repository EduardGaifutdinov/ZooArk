package request

// AddProduct request scheme
type AddClothes struct {
	Name  string
	Count int
	Price float64
	Type  string
	Color string
	Stock string
} // @name AddClothesRequest

// Delete request scheme
type DeleteClothes struct {
	Count int
} // @name DeleteClothesRequest