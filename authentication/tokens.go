package authentication

import (
	"ORM/commons"
	"ORM/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user models.LoginUser, w http.ResponseWriter) string {
	// Create a new token object, specifying signing method and the claims you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": user.Email,
		"nbf": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("CLAVESECRETATOKEN")))
	if err != nil {
		commons.SendResponse(w, http.StatusNotAcceptable, "Problemas al generar el Token "+err.Error())
		return ""
	} else {
		return tokenString
	}
}

func ValidateToken(w http.ResponseWriter, r *http.Request) bool {
	//recupero la cookie

	tokenCookie, err := r.Cookie("AuthAppORM")
	tokenString := tokenCookie.String()

	if err != nil {
		return false
	}

	fmt.Println(tokenString)
	return true
}
