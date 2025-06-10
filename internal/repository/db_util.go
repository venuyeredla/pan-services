package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db_con_pool *sql.DB // Database connection pool.
//var isPool

const (
	DATA_BASE_NAME       = "mysql"
	DB_CONNECTION_STRING = "root:ecompwd#24@tcp(127.0.0.1:3306)/ecom" //192.168.68.101
	C_CREATE             = "CREATE TABLE Customer(id INTEGER PRIMARY KEY, name TEXT, price INT);"
)

func IntializePool() *sql.DB {
	var db_error error
	db_con_pool, db_error = sql.Open(DATA_BASE_NAME, DB_CONNECTION_STRING)
	if db_error != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", db_error)
	}
	db_con_pool.SetConnMaxLifetime(0)
	db_con_pool.SetMaxIdleConns(3)
	db_con_pool.SetMaxOpenConns(3)
	if err := db_con_pool.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Connected sucessfully")
	}
	return db_con_pool
}

func ClosePool() {
	db_con_pool.Close()
	fmt.Println("Closed connection pool sucessfully")
}

func ExecuteQuery(query string) (bool, error) {
	ctx := context.Background()
	sqlConn, con_err := db_con_pool.Conn(ctx)
	if con_err != nil {
		log.Fatal("Exceptin happned in getting connection")
	}
	defer sqlConn.Close()
	tx, tx_error := sqlConn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})
	tx.Commit()
	if tx_error != nil {
		log.Fatal("Error in initialzing transaction")
	}
	result, dbError := sqlConn.ExecContext(ctx, query)
	if dbError != nil {
		log.Fatal("Error in executing ")
		return false, dbError
	}
	rowsEffected, _ := result.RowsAffected()
	fmt.Printf("Rows effected = %v", rowsEffected)

	return true, nil
}
