package domain

type User struct {
	ID       string
	Username string
	Password string
}

type Order struct {
	ID        string
	OrderName string
	Price     float64
}
