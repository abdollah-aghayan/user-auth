package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type AuthDetails struct {
	UserId string
}

// CreateToken create a jwt token based on secret and expire after ex (unix time)
func CreateToken(authD AuthDetails, secret string, ex int64) (string, error) {
	// create a claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = authD.UserId
	claims["exp"] = ex

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func TokenValid(token *jwt.Token) (jwt.MapClaims, bool) {
	//the token claims should conform to MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, false
	}

	return claims, true
}

// VerifyToken parse and validate a token based on secret string
func VerifyToken(tokenString, secret string) (*jwt.Token, error) {

	// parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// check if the signing method is not SigningMethodHMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

//GetToken get the token from the request
func GetToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	// extract token Authorization "Berar token"
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenAuth(r *http.Request, secret string) (*AuthDetails, error) {
	reqToken := GetToken(r)
	token, err := VerifyToken(reqToken, secret)
	if err != nil {
		return nil, err
	}

	claims, ok := TokenValid(token)
	if !ok {
		return nil, errors.Errorf("Invalid token %v", token)
	}

	return &AuthDetails{
		UserId: fmt.Sprint(claims["user_id"]),
	}, nil

}
