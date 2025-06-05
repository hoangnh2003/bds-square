package request

type ProductUpdateRequest struct {
	Name  string `json:"name" example:"Coca Chai"`
	Price string `json:"price" example:"12000"`
}
