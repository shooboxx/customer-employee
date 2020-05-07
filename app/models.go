package app

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

type Store struct {
	customers []string `json:"customers"`
	orders    []string `json:"orders"`
	products  []string `json:"products"`
	employees []string `json:"employees"`
}

type Customer struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postal_code"`
	Country     string `json:"country"`
	CreditLimit string `json:"credit_limit"`
}

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json"description"`
	Quantity    string `json:"quantity"`
	BuyPrice    string `json:"buy_price"`
}

type Employee struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	JobTitle  string `json:"job_title"`
}

type Order struct {
	ID         int64  `json:"id"`
	CustomerId int64  `json:"customer_id"`
	EmployeeId int64  `json:"employee_id"`
	ProductId  int64  `json:"product_id"`
	Status     string `json:"status"`
}
