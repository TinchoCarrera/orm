package models

import (
	jwt "github.com/dgrijalva/jwt-go" //Importo y le agrego un alias
)

type Claims struct {
	User               `json:"user"`
	jwt.StandardClaims //Agrega los campos de Fecha de expiración, objetivo final, audiencia, etc.
}
