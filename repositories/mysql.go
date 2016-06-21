package repositories

import(
	"fmt"
	"log"
  _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

// TODO add dynamic query
func (repo *sqlRepository) GetAll(v interface{}) error{
	var (
			db *sql.DB
			rows *sql.Rows
		)
  	db, _ = sql.Open(driver, repo.user + ":" + repo.password + "@tcp(" + repo.host + ":" + repo.port + ")/" + repo.database)
	rows, err := db.Query("SELECT * FROM users")
    checkErr(err)

    for rows.Next() {
        var id int
        var name string
        var gender string
        var age int
        err = rows.Scan(&id, &name, &gender, &age)
        checkErr(err)
        fmt.Println(id)
        fmt.Println(name)
        fmt.Println(gender)
        fmt.Println(age)
    }
	return nil

}

// TODO add dynamic query
func (repo *sqlRepository) Create(interface{}) error{
	var (
			db *sql.DB
			stmt *sql.Stmt
		)
	 db, _ = sql.Open(driver, repo.user + ":" + repo.password + "@tcp(" + repo.host + ":" + repo.port + ")/" + repo.database)

	stmt, err := db.Prepare("INSERT users SET name=?,gender=?,age=?")
	if err != nil {
		 panic(err)
	}

	res, err := stmt.Exec("joe", "male", 30)
	if err != nil {
		 panic(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		 panic(err)
	}

	fmt.Println(id)
	return nil
}

// TODO add dynamic query
func (repo *sqlRepository) Get(id int, v interface{}) error{
	var (
		err error
		db *sql.DB
		stmt *sql.Stmt
		)
  db, _ = sql.Open(driver, repo.user + ":" + repo.password + "@tcp(" + repo.host + ":" + repo.port + ")/" + repo.database)

  stmt, err = db.Prepare("select * from users where id = ?")
  if err != nil {
  	log.Fatal(err)
  }
  var name string
  var gender string
  var age int
  err = stmt.QueryRow(id).Scan(&id, &name, &gender, &age)
  if err != nil {
  	log.Fatal(err)
  }
  fmt.Println(name)

db.Close()
	return nil

}

// TODO add dynamic query
func (repo *sqlRepository) Update(id int, v interface{}) error {
	var (
	 	err error
 		db *sql.DB
	  	stmt *sql.Stmt
  	  )
    db, _ = sql.Open(driver, repo.user + ":" + repo.password + "@tcp(" + repo.host + ":" + repo.port + ")/" + repo.database)
	stmt, err = db.Prepare("update users set name=? where id=?")
    checkErr(err)

    _, err = stmt.Exec("joe", id)
    checkErr(err)

    // affect, err := res.RowsAffected()
    // checkErr(err)

	return nil
}

// TODO add dynamic query
func (repo *sqlRepository) Delete(id int, v interface{}) error{
	var (
		err error
		db *sql.DB
		stmt *sql.Stmt
  	  )
     db, _ = sql.Open(driver, repo.user + ":" + repo.password + "@tcp(" + repo.host + ":" + repo.port + ")/" + repo.database)
	stmt, err = db.Prepare("delete from users where id=?")
    checkErr(err)

    res, err := stmt.Exec(id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()
	return nil
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
