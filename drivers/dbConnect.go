package cloudSql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DbUser = "root"      //os.Getenv("SQL_USER")
	DbPwd  = ""          //os.Getenv("SQL_PASSWORD")
	DbHost = "localhost" //os.Getenv("SQL_HOST")
	DbName = "crud"      //os.Getenv("SQL_DB")
)

type SqlClient struct {
	DB *sql.DB
}

func ConnectDB() *SqlClient {
	DbPort := "3306"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPwd, DbHost, DbPort, DbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return &SqlClient{
		DB: db,
	}
}
