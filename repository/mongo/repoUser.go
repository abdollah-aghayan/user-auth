package mongo_db

import (
	"context"
	"log"
	"os"
	"user-auth/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	db *mongo.Client
}

func (repo *userRepo) Close() {
	// close connection and ignore error
	err := repo.db.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func (repo *userRepo) getCollection() *mongo.Collection {
	return repo.db.Database("quiz").Collection("users")
}

// New return new user repo
func New(connString string) *userRepo {
	// Set client options
	clientOptions := options.Client().ApplyURI(connString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		os.Exit(2)

	}

	return &userRepo{
		db: client,
	}
}

// FindByEmail find by email
func (repo *userRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	// define filter
	filter := map[string]string{"email": email}

	// user obj
	res := domain.User{}

	// fetch user
	err := repo.getCollection().FindOne(ctx, filter).Decode(&res)

	// check if result is empty
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	// check error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateUser insert a user to the database
func (repo *userRepo) CreateUser(ctx context.Context, user domain.User) error {
	_, err := repo.getCollection().InsertOne(context.TODO(), user)
	return err
}

//FindByID find by id
func (repo *userRepo) FindByID(ctx context.Context, userID string) (*domain.User, error) {
	// define filter
	filter := map[string]string{"_id": userID}

	// user obj
	res := domain.User{}

	// fetch user
	err := repo.getCollection().FindOne(ctx, filter).Decode(&res)

	// check if result is empty
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	// check error
	if err != nil {
		return nil, err
	}

	return &res, nil
}
