package request

type ProductCreateRequest struct {
	Name  string `json:"name" example:"Coca Cola"`
	Price string `json:"price" example:"10000"`
}
