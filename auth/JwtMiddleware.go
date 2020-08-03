package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

// JwtMiddleware Creates the middleware
func JwtMiddleware() *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Check audience
			aud := os.Getenv("GUSHKIN_AUDIENCE")
			claims := token.Claims.(jwt.MapClaims)
			checkAud := claims.VerifyAudience(aud, false)
			if !checkAud {
				return token, errors.New("invalid audience")
			}

			// Check issuer
			iss := os.Getenv("GUSHKIN_ISSUER")
			checkIss := claims.VerifyIssuer(iss, false)
			if !checkIss {
				fmt.Println("Issuer issue")
				return token, errors.New("invalid issuer")
			}

			if claims.Valid() != nil {
				panic("Invalid Claims")
			}

			cert, err := getPemCert(token)

			if err != nil {
				panic(err.Error())
			}

			result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))

			if err != nil {
				fmt.Println("Parse error")
			}

			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
	})
}

func getPemCert(token *jwt.Token) (string, error) {
	var cert string

	resp, err := http.Get(os.Getenv("GUSHKIN_JWKS_URL"))

	if err != nil {
		fmt.Println("Error 64")
		return cert, err
	}

	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		fmt.Println("Error 74")
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")
		return cert, err
	}

	return cert, nil
}
