package mysql_db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"user-auth/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	queryInsertUser  = "INSERT INTO users (id, email, username, password) VALUES (?, ?, ?, ?)"
	queryFindByEmail = "SELECT * FROM users WHERE email = ?"
	queryFindById    = "SELECT * FROM users WHERE id = ?"
)

type userRepo struct {
	db *sqlx.DB
}

func (repo *userRepo) Close() {
	// close connection and ignore error
	err := repo.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// New return new user repo
func New(dbConnString string) *userRepo {
	//connect to db
	db, err := sqlx.Connect("mysql", dbConnString)

	if err != nil {
		fmt.Println("Can not connect to database ", err)
		os.Exit(2)
	}

	return &userRepo{db: db}
}

// FindByEmail find by email
func (repo *userRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	// prepare a statement
	stat, err := repo.db.Preparex(queryFindByEmail)

	if err != nil {
		log.Println("failed to prepar query", err)
		return nil, errors.New("internal server error")
	}
	defer stat.Close()

	// Fetch user
	var user domain.User
	err = stat.Get(&user, email)

	if err != nil {
		// check weather response empty or error rased
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, nil
		}
		log.Println("database error", err)
		return nil, errors.New("internal server error")
	}

	return &user, nil
}

// CreateUser insert a user to the database
func (repo *userRepo) CreateUser(ctx context.Context, user domain.User) error {
	// prepare a statement
	stat, err := repo.db.Prepare(queryInsertUser)

	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	defer stat.Close()

	// execute query
	_, err = stat.Exec(user.ID, user.Email, user.Username, user.Password)

	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}

	return nil
}

//FindByID find by id
func (repo *userRepo) FindByID(ctx context.Context, userID string) (*domain.User, error) {
	// prepare query statement
	stat, err := repo.db.Preparex(queryFindById)

	if err != nil {
		log.Println("failed to prepar query", err)
		return nil, errors.New("internal server error")
	}
	defer stat.Close()

	// Fetch user
	var user domain.User
	err = stat.Get(&user, userID)

	if err != nil {
		// check weather response empty or error rased
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, nil
		}

		// return error
		log.Println("database error", err)
		return nil, errors.New("internal server error")
	}

	return &user, nil
}
