package security

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ribeirosaimon/go_flight_api/src/model"
	"github.com/ribeirosaimon/go_flight_api/src/repository"
	"log"
	"time"
)

var _SECRET_KEY = ""

func init() {
	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	_SECRET_KEY = base64.StdEncoding.EncodeToString(key)
}

func CreateToken(account model.Account) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 24).Unix()
	permission["userId"] = account.ID.Hex()
	return jwt.NewWithClaims(jwt.SigningMethodHS256, permission).SignedString([]byte(_SECRET_KEY))
}

func ValidationToken(token string) (model.LoggedUser, error) {
	parseToken, err := jwt.Parse(token, verifyKey)
	if err != nil {
		return model.LoggedUser{}, err
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)

	if ok && parseToken.Valid {
		userId := claims["userId"]
		userDb, err := repository.FindById(fmt.Sprint(userId))
		if err != nil {
			return model.LoggedUser{}, err
		}
		return model.LoggedUser{
			Username: userDb.Username,
			UserId:   userDb.ID.Hex(),
			Roles:    userDb.Roles,
		}, nil
	}

	return model.LoggedUser{}, errors.New("invalid Token")
}

func verifyKey(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("erro in token method")
	}
	return []byte(_SECRET_KEY), nil
}
