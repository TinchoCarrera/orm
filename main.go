package main

import (
	"ORM/commons"
	"ORM/db"
	initializars "ORM/initializers"
	"ORM/models"
	"ORM/routes"
)

func init() {
	//Conecto las variables de Entorno
	initializars.LoadEnvVariables()

	//Conecto la Base de Datos
	db.DBConnection()

	//Creo y actualizo las bases de datos
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	//Activo el cron
	commons.SetCron()

}

func main() {

	//Defino las rutas e inicio el server en 8000
	routes.StartRoutes()

	//
}
