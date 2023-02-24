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
	/*
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return os.Getenv("CLAVESECRETATOKEN"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
			return true
		} else {
			fmt.Println(err.Error())
			return false
		}
	*/

}
