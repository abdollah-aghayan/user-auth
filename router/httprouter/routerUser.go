package httprouter

import (
	"fmt"
	"net/http"
	"time"
	"user-auth/config"
	"user-auth/domain"
	"user-auth/logic"
	"user-auth/utils/auth"
	"user-auth/utils/errorh"

	"github.com/gin-gonic/gin"
)

//UserRequest user request interface
type UserRequest interface {
	getUserInfo(c *gin.Context)
	registerUser(c *gin.Context)
	loginUser(c *gin.Context)
}

//userRoute
type userRoute struct {
	UserRequest
}

var userRoutes = newUserRoute()

func newUserRoute() UserRequest {
	return &userRoute{}
}

/**
 * Routes Method
 */

//getUserInfo get user info by id
func (u *userRoute) getUserInfo(c *gin.Context) {

	userId, exist := c.Get("user_id")
	if !exist {
		err := errorh.NotAuthorizedError("Not Authorized")
		c.JSON(err.Code, err)
	}

	userLogic := logic.NewUser()
	user, errh := userLogic.GetUserInfo(c, fmt.Sprint(userId))

	if errh != nil {
		c.JSON(errh.Code, errh)
		return
	}
	// remove password
	user.Password = ""

	c.JSON(200, user)
	return
}

// register user
func (u *userRoute) registerUser(c *gin.Context) {
	regUser := domain.RegisterUser{}
	c.BindJSON(&regUser)

	// validate data
	err := regUser.ValidateRegisterUser()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	// call logic
	userLogic := logic.NewUser()
	userInfo, err := userLogic.RegisterUser(c, regUser)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	// remove password
	userInfo.Password = ""

	// response
	c.JSON(http.StatusCreated, userInfo)
	return
}

// login user
func (u *userRoute) loginUser(c *gin.Context) {

	loginUsr := domain.LoginUser{}

	c.BindJSON(&loginUsr)

	err := loginUsr.ValidateLoginUser()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	usrLogic := logic.NewUser()
	user, err := usrLogic.LoginUser(c, loginUsr)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	// set password to empty value to remove it form response
	user.Password = ""

	// expire time
	ex := time.Now().Add(10 * time.Minute).Unix()
	// token data
	authDetail := auth.AuthDetails{
		UserId: user.ID,
	}

	// create token
	token, _ := auth.CreateToken(authDetail, config.SECERET, ex)

	res := map[string]interface{}{
		"user":  user,
		"token": token,
	}

	c.JSON(http.StatusOK, res)
	return

}
