package repo

import "fmt"

type Mongo struct{}

func NewMongo() *Mongo {
	return &Mongo{}
}

func (m *Mongo) insert() {
	fmt.Println("inserted...")
}
