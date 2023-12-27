package main

import "dpadapter/repo"

func main() {
	mongoAdapter := repo.NewMongoAdapter()
	mongoAdapter.InsertOne()
}
