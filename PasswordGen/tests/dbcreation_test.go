package tests

import (
	"testing"

	db "passwordgen/dbconnection"
)

const (
	username  = "root"
	host_port = "127.0.0.1:3306"
	password  = "mynewpassword"
	dbname    = "test"

	name = "BobsWebsite"
	url = "www.bobsrecipes.com"
	uname = "bobsusername"
	pword = "password1"
	
)


func TestDBConnection(t *testing.T) {
	db.CreateDB(dbname)
}

func TestOpenDB(t *testing.T) {
	db.OpenDB(dbname)
}

func TestInsertIntoDB(t *testing.T){
	db.CreateDB(dbname)
	newdb, _ := db.OpenDB(dbname)

	db.CreateTables(newdb)
	db.InsertIntoDB(name,url,uname,pword,newdb)
	
}

