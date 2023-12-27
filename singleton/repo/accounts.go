package repo

type AccountsRepository struct{}

func NewAccountsRepository() *AccountsRepository {
	return &AccountsRepository{}
}

func (c *AccountsRepository) ListAll() []string {
	return []string{"Account1", "Account2"}
}
