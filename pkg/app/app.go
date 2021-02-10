package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niconosenzo/devopsapi/pkg/app/handler"
	"github.com/niconosenzo/devopsapi/pkg/app/model"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	Users  []model.User
}

//Initialize App with predefined users & configuration
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
	a.Users = []model.User{
		{ID: "1", Name: "Jose", Surname: "Perez"},
		{ID: "2", Name: "Pablo", Surname: "Martinez"},
	}
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Post("/user", a.CreateUser)
	a.Get("/user/{id}", a.GetUser)
	a.Get("/users", a.GetAllUsers)
	a.Delete("/user/{id}", a.DeleteUser)
}

//Get Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

//Post Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

//Delete Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	a.Users = handler.CreateUser(w, r, a.Users)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	handler.GetUser(w, r, a.Users)
}

func (a *App) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllUsers(w, r, a.Users)
}

func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	a.Users = handler.DeleteUser(w, r, a.Users)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
