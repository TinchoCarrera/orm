package controllers

import (
	"ORM/authentication"
	"ORM/commons"
	"ORM/db"
	"ORM/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	if authentication.ValidateToken(w, r) == false {
		commons.SendResponse(w, http.StatusNotFound, "Usuario sin permisos")
		return

	}

	var users []models.User

	db.DB.Find(&users)

	//Envio el resultado
	commons.SendResponse(w, http.StatusOK, users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	//Guardo en vars todo lo que viene en la URL por GET
	params := mux.Vars(r)

	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		commons.SendResponse(w, http.StatusNotFound, "El Usuario no existe")
		return
	}

	//agrego todas las tareas que el usuario tiene
	db.DB.Model(&user).Association("Task").Find(&user.Task)

	//Envio el resultado
	commons.SendResponse(w, http.StatusOK, user)

}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	//guardo en user todo lo que viene por POSTs
	json.NewDecoder(r.Body).Decode(&user)

	//hasheo el password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		commons.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	} else {
		user.Password = string(hash)
	}

	//Creo un usuario
	createdUser := db.DB.Create(&user)
	err = createdUser.Error
	if err != nil {
		commons.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//Envio el resultado
	commons.SendResponse(w, http.StatusCreated, user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//Guardo en vars todo lo que viene en la URL por GET
	params := mux.Vars(r)

	var user models.User
	db.DB.First(&user, params["id"])

	//agrego todas las tareas que el usuario tiene
	db.DB.Model(&user).Association("Task").Find(&user.Task)

	//borro el usuario encontrado pero lo deje en la tabla con fecha de eliminación
	db.DB.Delete(&user)
	//borro el usuario encontrado de la tabla
	//db.DB.Unscoped().Delete(&user)

	if user.ID == 0 {
		commons.SendResponse(w, http.StatusNotFound, "El Usuario no existe")
		return
	}

	//Envio el resultado
	commons.SendResponse(w, http.StatusOK, user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user, newUser models.User
	//guardo en user todo lo que viene por POSTs
	json.NewDecoder(r.Body).Decode(&newUser)

	//busco el usuario
	//Guardo en vars todo lo que viene en la URL por GET
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	fmt.Println(newUser.FirstName)
	user.FirstName = newUser.FirstName
	user.LastName = newUser.LastName
	user.Email = newUser.Email
	db.DB.Updates(&user)

	//Envio el resultado
	commons.SendResponse(w, http.StatusOK, user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser models.LoginUser
	//guardo en user todo lo que viene por POSTs
	json.NewDecoder(r.Body).Decode(&loginUser)

	//busco el user por mail
	var user models.User
	//db.DB.Find(&user, "email='"+loginUser.Email+"'")
	db.DB.Find(&user, "email", loginUser.Email)

	if user.ID == 0 {
		commons.SendResponse(w, http.StatusNotFound, "El Usuario no existe")
		return
	}

	//Comparo los hashes
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	fmt.Println("err:", err)
	if err != nil {
		commons.SendResponse(w, http.StatusUnauthorized, "Las contraseñas no coinciden")
		return
	}

	//Creo el token de seguridad
	tokenString := authentication.CreateToken(loginUser, w)

	if tokenString != "" {
		//Creamos la cookie
		authentication.CreateCookie(tokenString, w)

		//envío el resultado
		commons.SendResponse(w, http.StatusOK, user)

	}

}
