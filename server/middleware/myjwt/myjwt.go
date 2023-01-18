package myjwt

import (
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jayanthkrishna/golang-csrf-project/db"
	"github.com/jayanthkrishna/golang-csrf-project/db/models"
)

const (
	privKeypath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

func InitJWT() error {
	signBytes, err := ioutil.ReadFile(privKeypath)
	if err != nil {
		return err
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	if err != nil {
		return err
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)

	if err != nil {
		return err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

	if err != nil {
		return err
	}

	return nil
}

func CreateNewTokens(uuid, role string) (authTokenString, refreshTokenString) {

	//generate the csrf secret

	csrfSecret, err := models.GenerateCSRFSecret()

	if err != nil {
		return
	}

	//create refresh token string

	refreshToken, err := createRefreshTokenString(uuid, role, csrfSecret)
	authToken, err := createAuthTokenString(uuid, role, csrfSecret)

	if err != nil {
		return nil, nil
	}

	return
}

func CheckAndRefreshTokens() {

}

func createAuthTokenString(uuid string, role string, csrfSecret string) (authTokenString string, err error) {
	authTokenExp := time.Now().Add(models.AuthTokenValidTime).Unix()
	authClaims := models.TokenClaims{

		jwt.StandardClaims{
			Subject:   uuid,
			ExpiresAt: authTokenExp,
		},
		role,
		csrfSecret,
	}

	authJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), authClaims)

	authTokenString, err = authJwt.SignedString(signKey)
	return authTokenString, err

}

func createRefreshTokenString(uuid string, role string, csrfSecret string) (refreshTokenString string, err error) {
	refreshTokenExp := time.Now().Add(models.RefreshTokenValidTime).Unix()
	refreshJti, err := db.StoreRefreshToken()

	if err != nil {
		return
	}

	refreshClaims := models.TokenClaims{
		jwt.StandardClaims{
			Id:        refreshJti,
			Subject:   uuid,
			ExpiresAt: refreshTokenExp,
		},
		role,
		csrfSecret,
	}

	refreshJwt := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), refreshClaims)
	refreshTokenString, err = refreshJwt.SigningString(signKey)

	return
}

func updateRefreshTokenString() {

}

func updateAuthTokenString() {

}

func RevokeRefreshToken(value string) {

}

func updateRefreshTokenCsrf() {

}

func GrabUUID() {

}
