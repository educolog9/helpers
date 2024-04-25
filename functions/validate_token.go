package functions

import (
	"net/http"
	"os"
	"strings"

	custom_errors "github.com/educolog9/helpers/errors"
	"github.com/educolog9/helpers/types"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(r *http.Request, w http.ResponseWriter, header string) *types.UserClaims {
	authHeader := r.Header.Get(header)
	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) != 2 {
		custom_errors.ReturnUnauthorizedError(w, []string{"Invalid token"})
		return nil
	}

	tokenString := bearerToken[1]
	token, err := jwt.ParseWithClaims(tokenString, &types.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		custom_errors.ReturnUnauthorizedError(w, []string{"Invalid or expired token"})
		return nil
	}

	if claims, ok := token.Claims.(*types.UserClaims); ok && token.Valid {
		return claims
	} else {
		custom_errors.ReturnUnauthorizedError(w, []string{"Invalid token"})
		return nil
	}
}
