package model

type CreateOrderInput struct {
	UserID string            `json:"userId"`
	Items  []*OrderItemInput `json:"items"`
}

type Order struct {
	ID         string       `json:"id"`
	User       *User        `json:"user"`
	Status     string       `json:"status"`
	TotalPrice float64      `json:"totalPrice"`
	Items      []*OrderItem `json:"items,omitempty"`
	CreatedAt  string       `json:"createdAt"`
	UpdatedAt  string       `json:"updatedAt"`
}

type OrderItem struct {
	ID        string   `json:"id"`
	Order     *Order   `json:"order"`
	Product   *Product `json:"product"`
	Quantity  int      `json:"quantity"`
	Price     float64  `json:"price"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

type OrderItemInput struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type UpdateOrderInput struct {
	Status *string `json:"status,omitempty"`
}
