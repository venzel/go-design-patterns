package repo

type MongoAdapter struct {
	repo RepoInterface
}

func NewMongoAdapter() *MongoAdapter {
	repo := &Mongo{}
	return &MongoAdapter{repo: repo}
}

func (m *MongoAdapter) InsertOne() {
	m.repo.insert()
}
