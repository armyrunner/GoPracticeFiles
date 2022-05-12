package dbconnection

import (
	"database/sql"
	"fmt"
	"strconv"



)

func check(e error){
  if e != nil{
    panic(e)
  }

}

func DBConn(host string, port int ,username ,password ,dbname string)(string,error){

	// connStr := "user="+ user + " dbname="+ dbname + " password=" + password + " host="+ host + " sslmode=disable"
  newport := strconv.Itoa(port)
  connStr := username + ":"+password+"@"+"tcp("+host+":"+newport+")/"+dbname
	
	return connStr, nil

  }

  func OpenDB(connStr string)(*sql.DB,error){
    db, err := sql.Open("postgres",connStr)
    check(err)

    defer db.Close()

    fmt.Printf("\nSuccessfully connected to database!\n")
    return db,nil

  }

  func InsertIntoDB(Name,URL,Username,Password string)(*sql.DB,error){

    
    
  }



