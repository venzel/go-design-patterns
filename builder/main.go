package main

import (
	"dpbuilder/configs"
	"fmt"
)

func main() {
	qb := configs.NewQueueBuilder()

	qb.SetProducer("Almeida").SetConsumer("Tiago").SetExchange("Ex").Build()

	fmt.Println(qb.Producer)
	fmt.Println(qb.Consumer)
	fmt.Println(qb.Exchange)
}
