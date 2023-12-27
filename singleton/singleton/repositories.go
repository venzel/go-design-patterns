package singleton

import (
	"dpsingleton/repo"
	"sync"
)

var instance *RepositoriesSingleton
var once sync.Once

type RepositoriesSingleton struct {
	CustomersRepository *repo.CustomersRepository
	AccountsRepository  *repo.AccountsRepository
}

func GetInstance() *RepositoriesSingleton {
	once.Do(func() {
		instance = &RepositoriesSingleton{
			CustomersRepository: repo.NewCustomersRepository(),
			AccountsRepository:  repo.NewAccountsRepository(),
		}
	})

	return instance
}
