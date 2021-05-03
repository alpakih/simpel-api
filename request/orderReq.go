package request

type (
	CreateOrder struct {
		CustomerName string  `json:"customerName"`
		OrderAt      string  `json:"orderedAt"`
		Items        []items `json:"items"`
	}

	UpdateOrder struct {
		OrderID      uint    `json:"orderId"`
		CustomerName string  `json:"customerName"`
		OrderAt      string  `json:"orderedAt"`
		Items        []items `json:"items"`
	}

	items struct {
		ID          uint   `json:"lineItemId,omitempty"`
		ItemCode    string `json:"itemCode"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}
)
