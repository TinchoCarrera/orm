package main

import (
	"ORM/commons"
	"ORM/db"
	"ORM/models"
	"ORM/routes"
)

func main() {

	db.DBConnection()

	//Creo y actualizo las bases de datos
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	commons.SetCron()
	//Defino las rutas e inicio el server en 8000
	routes.StartRoutes()

	//
}
