package main

import (
	"flag"
	"fmt"
	"user-auth/config"
	"user-auth/repository"
	"user-auth/router/httprouter"
)

func main() {

	var dbType = flag.String("db_type", "mongodb", "The database to app work with")

	flag.Parse()

	// load config
	config.Init()

	// repository
	var connStr string
	if *dbType == repository.MOGODB {
		connStr = fmt.Sprintf("mongodb://%s:%s", config.MDB_HOST, config.MDB_PORT)
	} else if *dbType == repository.SQL {
		connStr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	} else {
		connStr = fmt.Sprintf("mongodb://%s:%s", config.MDB_HOST, config.MDB_PORT)
	}

	repo := repository.Init(*dbType, connStr)
	defer repo.Close()

	// router
	httprouter.Run(config.HTTPPort)
}
