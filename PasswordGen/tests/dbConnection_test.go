package tests

import (
	"fmt"
	"testing"

	db "passwordgen/dbconnection"
)

func TestDBConnectionStr(t *testing.T) {

	ConnStr := []struct {
		host     string
		port     int
		username string
		password string
		dbname   string
	}{
		{host: "localhost", port: 5432, username: "postgres", password: "secure-password", dbname: "connect-db"},
	}

	for _, dbTestStr := range ConnStr {
		constr, _ := db.DBConn(dbTestStr.host, dbTestStr.port, dbTestStr.username, dbTestStr.password, dbTestStr.dbname)
		fmt.Printf("constr: %v\n", constr)
	}

}
