package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/apex/log"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

// init loads the private and public keys from the files
func init() {
	fmt.Println("This is the initialization")

	// this location is relative to the main file
	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("Error reading private key file")
	}

	// this location is relative to the main file
	publicBytes, err := ioutil.ReadFile("./public.rsa.pub")
	if err != nil {
		log.Fatal("Error reading public key file")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Error parsing private key")
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Error parsing private key")
	}
}

// GenerateJWT creates a JWT using the user info received.
// Returns the token already in base64 encoding
func GenerateJWT(user string) string {
	claims := Claim{
		UserID: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // expires one hour from now in unix time
			Issuer:    "issuer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenBase64, err := token.SignedString(privateKey)
	if err != nil {
		log.Error("Error signing token string")
	}

	return tokenBase64
}
