package dbconnection

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

func check(e error){
  if e != nil{
    panic(e)
  }

}

func PostgresConnStr(host string, port int ,username ,password ,dbname string)(string,error){

	// connStr := "user="+ user + " dbname="+ dbname + " password=" + password + " host="+ host + " sslmode=disable"
  newport := strconv.Itoa(port)
  connStr := username + ":"+password+"@"+"tcp("+host+":"+newport+")/"+dbname
	
	return connStr, nil

  }

  func OpenDB(connStr string){
    db, err := sql.Open("postgres",connStr)
    check(err)

    defer db.Close()

    err = db.Ping()
    check(err)

    fmt.Printf("\nSuccessfully connected to database!\n")

  }

  