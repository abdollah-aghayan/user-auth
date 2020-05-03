package logic

import (
	"context"
	"log"
	"user-auth/domain"
	"user-auth/utils/errorh"
	"user-auth/utils/passwordh"

	"github.com/google/uuid"
)

//UserLogicInterface interface
type UserLogicInterface interface {
	LoginUser(ctx context.Context, user domain.LoginUser) (*domain.User, *errorh.Errorh)
	RegisterUser(ctx context.Context, user domain.RegisterUser) (*domain.User, *errorh.Errorh)
	GetUserInfo(ctx context.Context, userId string) (*domain.User, *errorh.Errorh)
}

//UserLogic user logic
type UserLogic struct {
	Self UserLogicInterface
	factory
}

//NewUserWithFactory create new user
// create user logic with inject factory
func NewUserWithFactory(sourceFactory factory) UserLogicInterface {
	x := UserLogic{
		factory: sourceFactory,
	}
	x.Self = &x
	return &x
}

//NewUser create new user
// use for production, because we don't need to dependency injection
func NewUser() UserLogicInterface {
	return NewUserWithFactory(NewFactory())
}

//LoginUser login user if credentials are correct
func (u *UserLogic) LoginUser(ctx context.Context, user domain.LoginUser) (*domain.User, *errorh.Errorh) {

	// get repo
	userRepo := u.factory.GetUserRepo(ctx)

	// check user if exist
	item, err := userRepo.FindByEmail(ctx, user.Email)
	if err != nil {
		log.Fatal(err)
		return nil, errorh.InternalError("Error raised try later")
	}

	if item == nil {
		return nil, errorh.NotFoundError("You don't have account")
	}

	if passwordh.ComparePasswords(item.Password, user.Password) {
		return item, nil
	}

	return nil, errorh.NotAuthorizedError("Wrong credential")
}

//GetUserByID get user by id
func (u *UserLogic) RegisterUser(ctx context.Context, user domain.RegisterUser) (*domain.User, *errorh.Errorh) {

	// check and inject
	userRepo := u.factory.GetUserRepo(ctx)

	// check if email exist
	eu, err := userRepo.FindByEmail(ctx, user.Email)
	if err != nil {
		log.Fatal(err)
		return nil, errorh.InternalError("Something went wrong please try later")
	}

	if eu != nil {
		return nil, errorh.BadRequestError("Email exist")
	}

	pwd, err := passwordh.CreatePassword(user.Password)

	if err != nil {
		log.Fatal(err)
		return nil, errorh.InternalError("Something went wrong please try later")
	}

	regUser := domain.User{
		ID:       uuid.New().String(),
		Email:    user.Email,
		Username: user.Username,
		Password: string(pwd),
	}

	// get user info
	err = userRepo.CreateUser(ctx, regUser)
	if err != nil {
		log.Fatal(err)
		return nil, errorh.InternalError("Something went wrong please try later")
	}

	return &regUser, nil
}

// GetUserInfo fetch user info based on userId
func (u *UserLogic) GetUserInfo(ctx context.Context, userId string) (*domain.User, *errorh.Errorh) {
	// get repo
	userRepo := u.factory.GetUserRepo(ctx)

	user, errh := userRepo.FindByID(ctx, userId)
	if errh != nil {
		log.Fatal(errh)
		return nil, errorh.InternalError("Somethinf went wrong try later")
	}

	if user == nil {
		return nil, errorh.NotFoundError("You don't have account")
	}

	return user, nil
}
