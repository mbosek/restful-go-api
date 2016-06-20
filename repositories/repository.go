package repositories

type (
	Repository interface {
		GetAll(interface{}) error
		Get(int, interface{}) error
		Create(interface{}) error
		Update(int, interface{}) error
		Delete(int, interface{}) error
	}

	repositories map[string]Repository
)

var (
	Repositories = make(repositories)
)

func Register(key string, r Repository) {
	Repositories[key] = r;
}
