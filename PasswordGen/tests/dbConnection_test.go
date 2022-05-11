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
		user     string
		password string
		dbname   string
	}{
		{host: "localhost", port: 5432, user: "postgres", password: "secure-password", dbname: "connect-db"},
	}

	for _, dbTestStr := range ConnStr {
		constr, _ := db.PostgresConnStr(dbTestStr.host, dbTestStr.port, dbTestStr.user, dbTestStr.password, dbTestStr.dbname)
		fmt.Printf("constr: %v\n", constr)
	}

}
