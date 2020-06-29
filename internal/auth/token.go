type ResponseToken struct {
	Token string `json:"token"`
}

// ValidateToken valites if the token is a valid one and haven't expired
func ValidateToken(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequestWithClaims(
		r,                       // takes the request token with all the claims included
		request.OAuth2Extractor, // type of extraction. OAuth2Extractor looks in 'Authorization' header
		&models.Claim{},         // struct or model of the claims to be extracted
		func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		},
	)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Fprintln(w, "Token expired")
				return
			case jwt.ValidationErrorSignatureInvalid:
				fmt.Fprintln(w, "Signature doesn't match")
				return
			default:
				fmt.Fprintln(w, "Invalid token")
				return
			}
		default:
			fmt.Fprintln(w, "Unknown token error")
			return
		}

	}

	if token.Valid {
		log.Println("is token valid?: ", token.Valid)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Token is no valid")
	}

	return
}