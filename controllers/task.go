package controllers

import (
	"ORM/commons"
	"ORM/db"
	"ORM/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)

	//Envio el resultado
	commons.SendResponse(w, http.StatusOK, tasks)

}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	//Guardo en vars todo lo que viene en la URL por GET
	params := mux.Vars(r)

	var task models.Task
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		commons.SendResponse(w, http.StatusNotFound, "La Tarea no existe")
		return
	}

	//Envio el resultado
	commons.SendResponse(w, http.StatusOK, task)

}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	//guardo en task todo lo que viene por POSTs
	json.NewDecoder(r.Body).Decode(&task)

	//Creo un usuario
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		commons.SendResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	//Envio el resultado
	commons.SendResponse(w, http.StatusCreated, task)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	//Guardo en vars todo lo que viene en la URL por GET
	params := mux.Vars(r)

	var task models.Task
	db.DB.First(&task, params["id"])
	//borro la tarea encontrada pero lo deje en la tabla con fecha de eliminación
	db.DB.Delete(&task)
	//borro la tarea encontrada de la tabla
	//db.DB.Unscoped().Delete(&task)

	if task.ID == 0 {
		commons.SendResponse(w, http.StatusNotFound, "La Tarea no existe")
		return
	}

	//Envio el resultado
	commons.SendResponse(w, http.StatusOK, task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task, newTask models.Task
	//guardo en task todo lo que viene por POSTs
	json.NewDecoder(r.Body).Decode(&newTask)

	//busco el usuario
	//Guardo en vars todo lo que viene en la URL por GET
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	task.Title = newTask.Title
	task.Description = newTask.Description
	task.Done = newTask.Done
	task.UserID = newTask.UserID

	db.DB.Updates(&task)

	//Envio el resultado
	commons.SendResponse(w, http.StatusOK, task)
}
