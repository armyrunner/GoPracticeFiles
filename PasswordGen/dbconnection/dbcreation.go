package dbconnection

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username  = "root"
	password  = "mynewpassword"
	host_port = "127.0.0.1:3306"
)

func DBConn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host_port, dbname)
}

func CreateDB(dbname string) (*sql.DB, error) {
	db, err := sql.Open("mysql", DBConn(""))
	if err != nil {
		log.Printf("Error in opening Database, %s", err)
		return nil, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating the database", err)
		return nil, err
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("rows affected %d\n", no)
	db.Close()

	return db, nil
}

func OpenDB(dbname string) (*sql.DB, error) {

	db, err := sql.Open("mysql", DBConn(dbname))
	if err != nil {
		log.Printf("Error %s when opening the database ", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Error %s pinging DB", err)
	}

	log.Printf("Connecting to %s was SUCCESSFUL!!!", dbname)

	return db, nil

}

func CreateTables(db *sql.DB) error {

	newTable := `CREATE TABLE IF NOT EXISTS useraccount(user_id int primary key auto_increment, name text, url text,
		username text, password text)`

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := db.ExecContext(ctx, newTable)
	if err != nil {
		return nil
	}

	return nil
}

func InsertIntoDB(name, url, uname, pword string, db *sql.DB) error {

	newTable := "INSERT INTO useraccount(name, url, username, password) VALUES(?,?,?,?)"

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	_, err := db.ExecContext(ctx, newTable, name, url, uname, pword)
	if err != nil {
		return nil
	}

	return nil
}

func UpdateIntoDB(name, pword string, db *sql.DB) error {

	newTable := "UPDATE useraccount password  = ? WHERE name = ?"

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	res, err := db.ExecContext(ctx, newTable, name, pword)
	if err != nil {
		log.Printf("Error %s when updating the database", err)
		return nil
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return err
	}

	log.Printf("Rows effected when updating the table %d", rows)

	return nil
}

func SelectingAUser(name string, db *sql.DB) (string, string, error) {

	newTable := "SELECT username,password FROM useraccount WHERE name = ?"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, newTable)
	if err != nil {
		log.Printf("error %s when preparing SQL Statement", err)
	}

	defer stmt.Close()

	var uname, pass string
	row := stmt.QueryRowContext(ctx, name)
	if err := row.Scan(&uname, &pass); err != nil {
		return "", "", err
	}

	return uname, pass, nil

}

func DeleteUserInfo(name string, db *sql.DB) error {
	deluser := "DELETE FROM useraccount WHERE name = ?"

	stmt, err := db.Prepare(deluser)
	if err != nil {
		log.Printf("error %s when preparing SQL Statement", err)
	}

	defer stmt.Close()


	_,err = stmt.Exec(name)
	if err != nil {
		return err
	}

	return nil
}
