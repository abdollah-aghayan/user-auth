package repository

import (
	"context"
	"fmt"
	"log"
	"user-auth/domain"

	mongo_db "user-auth/repository/mongo"
	mysql_db "user-auth/repository/mysql"
)

const (
	MOGODB = "mongodb"
	SQL    = "sql"
)

//UserRepositoryInterface interface
type UserRepositoryInterface interface {
	Close()
	FindByID(ctx context.Context, userID string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) error
}

var db UserRepositoryInterface

func Init(dbType, connString string) UserRepositoryInterface {
	fmt.Println("Connecting to db...")
	switch dbType {
	case MOGODB:
		// connect to mongodb
		db = mongo_db.New(connString)
	case SQL:
		// connect to mysql
		db = mysql_db.New(connString)
	default:
		// connect ot default db
		db = mongo_db.New(connString)

	}

	return db
}

func GetDB() UserRepositoryInterface {
	if db == nil {
		log.Fatal("Please init first")
		return nil
	}
	return db
}
