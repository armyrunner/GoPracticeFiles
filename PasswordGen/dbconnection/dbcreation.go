package dbconnection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func check(e error){
  if e != nil{
    panic(e)
  }

}

func PostgresConnStr(host string,port int,user ,password ,dbname string)(string,error){

	connStr := "user="+ user + " dbname="+ dbname + " password=" + password + " host="+ host + " sslmode=disable"
	
	return connStr, nil

  }

  func OpenDB(connStr string){
    db, err := sql.Open("postgres",connStr)
    check(err)

    defer db.Close()

    err = db.Ping()
    check(err)

    fmt.Printf()

  }