package repo

type CustomersRepository struct{}

func NewCustomersRepository() *CustomersRepository {
	return &CustomersRepository{}
}

func (c *CustomersRepository) ListAll() []string {
	return []string{"John", "Doe"}
}
