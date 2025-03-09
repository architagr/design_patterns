package models

type GetProductResponse struct {
	ProductId    int    `json:"productId"`
	ProductName  string `json:"productName"`
	ProductImage string `json:"productImage"`
}
