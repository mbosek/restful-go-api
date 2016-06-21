package repositories

type (
	sqlRepository struct {
		host     string
		port     string
		database string
		user     string
		password string
	}
	Repository interface {
		GetAll(interface{}) error
		Get(int, interface{}) error
		Create(interface{}) error
		Update(int, interface{}) error
		Delete(int, interface{}) error
	}

	repositories map[string]Repository
)

const driver string = "mysql"
var Repositories = make(repositories)

func NewRepository(database string, args ...string) Repository {
	return &sqlRepository{"192.168.99.100", "3306", database, "root", "password"}
}

func Register(key string, r Repository) {
	Repositories[key] = r;
}
