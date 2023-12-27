package configs

type QueueBuilder struct {
	Producer string
	Consumer string
	Exchange string
}

func NewQueueBuilder() *QueueBuilder {
	return &QueueBuilder{}
}

func (qb *QueueBuilder) SetProducer(producer string) *QueueBuilder {
	qb.Producer = producer
	return qb
}

func (qb *QueueBuilder) SetConsumer(consumer string) *QueueBuilder {
	qb.Consumer = consumer
	return qb
}

func (qb *QueueBuilder) SetExchange(exchange string) *QueueBuilder {
	qb.Exchange = exchange
	return qb
}

func (qb *QueueBuilder) Build() *QueueBuilder {
	return qb
}
