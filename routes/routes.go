package routes

import (
	"ORM/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartRoutes() {
	r := mux.NewRouter()

	//Ruta Principal.....
	r.HandleFunc("/", controllers.IndexRoute)

	//Rutas de Usuarios
	r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", controllers.UpdateUserHandler).Methods("PUT")

	//Rutas de Tareas
	r.HandleFunc("/tasks", controllers.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", controllers.UpdateTaskHandler).Methods("PUT")

	//login
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	//Inicio el Servidor
	http.ListenAndServe(":"+os.Getenv("PORT"), r)

	//esto es 100% nuevo

}
