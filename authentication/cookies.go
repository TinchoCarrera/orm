package authentication

import (
	"net/http"
	"time"
)

func CreateCookie(tokenString string, w http.ResponseWriter) {
	//Guardo el token en una cookie
	var cookie = http.Cookie{
		Name:       "AuthAppORM",
		Expires:    time.Now().Add(time.Hour * 24 * 30 * 5),
		RawExpires: "",
		Secure:     false,
		HttpOnly:   true,
		SameSite:   http.SameSiteStrictMode,
		Value:      tokenString,
	}

	http.SetCookie(w, &cookie)
}
