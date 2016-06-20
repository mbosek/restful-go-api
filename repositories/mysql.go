package repositories

import(
	"fmt"
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
)

// type dbConnection func() (db *sql.DB)

// var Db dbConnection
// func init(){
// 	Db = connection
// }

// func connection() (db *sql.DB) {
//
// 	dbDriver := "mysql"
// 	dbAddress := "192.168.99.100"
// 	dbPort := "3306"
// 	dbUser := "root"
// 	dbPass := "password"
// 	dbName := "go"
//
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbAddress+":"+dbPort+")"+"/"+dbName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

const driver string = "mysql"

type (
	sqlRepository struct {
		host     string
		port     string
		database string
		user     string
		password string
	}
)

func NewRepository(database string, args ...string) Repository {

	config := [4]string{"127.0.0.1", "3306", "root", "password"}
	for i, val := range args {
		config[i] = val
		if i == 3 {
			break;
		}
	}

	return &sqlRepository{config[0], config[1], database, config[2], config[3]}
}

func (repo *sqlRepository) GetAll(v interface{}) error{
	var (
			db *sql.DB
			rows *sql.Rows
		)
  db, _ = sql.Open(driver, repo.user + ":" + repo.password + "@tcp(" + repo.host + ":" + repo.port + ")/" + repo.database)
	rows, err := db.Query("SELECT * FROM users")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }
		return nil

}

func (repo *sqlRepository) Get(id int, v interface{}) error{
	fmt.Println("asdasd")
	return nil
}

func (repo *sqlRepository) Create(interface{}) error{
fmt.Println("asdasd")
return nil
	// db := connection()
	// stmt, err := db.Prepare("INSERT users SET username=?,departname=?,created=?")
	// if err != nil {
	// 	 panic(err)
	// }
	//
	// res, err := stmt.Exec("astaxie", "asasd", "2012-12-09")
	// if err != nil {
	// 	 panic(err)
	// }
	//
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	 panic(err)
	// }
	//
	// fmt.Println(id)
}

func (repo *sqlRepository) Update(id int, v interface{}) error{
	fmt.Println("asdasd")
	return nil
  // db := connection()
	// stmt, err = db.Prepare("update users set username=? where uid=?")
  //   checkErr(err)
	//
  //   res, err = stmt.Exec("astaxieupdate", id)
  //   checkErr(err)
	//
  //   affect, err := res.RowsAffected()
  //   checkErr(err)
	//
  //   fmt.Println(affect)
}

func (repo *sqlRepository) Delete(id int, v interface{}) error{
	fmt.Println("asdasd")
	return nil
  // db := connection()
	// stmt, err = db.Prepare("delete from users where uid=?")
  //   checkErr(err)
	//
  //   res, err = stmt.Exec(id)
  //   checkErr(err)
	//
  //   affect, err = res.RowsAffected()
  //   checkErr(err)
	//
  //   fmt.Println(affect)
	//
  //   db.Close()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
