package authentication

import (
	"ORM/commons"
	"ORM/models"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	Privatekey *rsa.PrivateKey
	Publickey  *rsa.PublicKey
)

// esta función se ejecuta al llamar al paquete
func init() {
	privatebytes, err := ioutil.ReadFile("authentication/private.rsa")
	if err != nil {
		log.Fatal("Error al leer el archivo Privado")
	}

	publicbytes, err := ioutil.ReadFile("authentication/public.rsa.pub")
	if err != nil {
		log.Fatal("Error al leer el archivo Público")
	}

	Privatekey, err = jwt.ParseRSAPrivateKeyFromPEM(privatebytes)
	if err != nil {
		log.Fatal("No se puede hacer el parse a Privatekey")
	}

	Publickey, err = jwt.ParseRSAPublicKeyFromPEM(publicbytes)
	if err != nil {
		log.Fatal("No se puede hacer el parse a Publickey")
	}
}

func GenerateJWToken(user models.User) string {
	claims := models.Claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Taller de Sábazdo",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	//Convierto el token en un string de base64
	result, err := token.SignedString(Privatekey)
	if err != nil {
		log.Fatal("No se puede firmar el token")
	}

	return result

}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Fprintln(w, "Error al leer el usuario: $v", err.Error())
		return
	}

	if user.Email == "tinchocarrera@gmail.com" && user.Password == "123" {
		user.Password = ""
		user.Role = "Admin"
		token := GenerateJWToken(user)

		response := models.ResponseToken{
			Token: token,
		}

		jsonresult, err := json.Marshal(response)

		if err != nil {
			fmt.Fprintln(w, "Error al leer el convertir a JSON el resultado: $v", err.Error())
			return
		}

		commons.SendResponse(w, http.StatusOK, jsonresult)
	} else {
		commons.SendResponse(w, http.StatusLocked, "Usuario y clave inválidos")
	}

}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Publickey, nil
	})

	if err != nil {
		//preguntamos por el tipo de error
		switch err.(type) {
		case *jwt.ValidationError:
			//Si es un error de validación de Token tenemos varias posibilidades
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				commons.SendResponse(w, http.StatusUnauthorized, "Su Token ha expirado")
				return
			case jwt.ValidationErrorSignatureInvalid:
				commons.SendResponse(w, http.StatusUnauthorized, "La Firma del token no coincide")
				return
			default:
				commons.SendResponse(w, http.StatusUnauthorized, "Su Token no es válido..."+vErr.Error())
				return
			}
		//si es otro tipo de error
		default:
			commons.SendResponse(w, http.StatusUnauthorized, "Existe algún tipo de error"+err.Error())
			return
		}
	}

	if token.Valid {
		commons.SendResponse(w, http.StatusOK, "Bienvenido al Sistema")
	} else {
		commons.SendResponse(w, http.StatusUnauthorized, "Su Token no es válido")
	}

}
