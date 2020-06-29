package auth

import jwt "github.com/dgrijalva/jwt-go"

// Claim contains the JWT payload information
type Claim struct {
	MemberID string `json:"member_id"`
	jwt.StandardClaims
}
