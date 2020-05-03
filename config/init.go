package config

import "fmt"

// only use for demo project
const (
	//HTTPPort http port
	HTTPPort = "6050"

	SECERET = "asdas721"

	DB_USER = "root"
	DB_PASS = "root"
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DB_NAME = "kingquiz"

	MDB_PORT = "27017"
	MDB_HOST = "localhost"

	DB_TYPE = "mongodb" // sql
)

//Init manual initialize
func Init() {
	fmt.Println("config loader...")
}
