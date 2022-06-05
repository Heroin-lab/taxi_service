package repositories

type Authorization interface {
}

type Repositories struct {
	Authorization
}

func NewRepositories() *Repositories {
	return &Repositories{}
}
