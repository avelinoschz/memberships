package auth

import jwt "github.com/dgrijalva/jwt-go"

// Claim contains the JWT payload information
type Claim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
