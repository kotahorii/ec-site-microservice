package model

type CreateOrderInput struct {
	UserID string            `json:"userId"`
	Items  []*OrderItemInput `json:"items"`
}

type Order struct {
	ID         string  `json:"id"`
	Status     string  `json:"status"`
	TotalPrice float64 `json:"totalPrice"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

type OrderItem struct {
	ID        string  `json:"id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type OrderItemInput struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type UpdateOrderInput struct {
	ID     string  `json:"id"`
	Status *string `json:"status,omitempty"`
}
