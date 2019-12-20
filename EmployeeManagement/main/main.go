package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	resthandlers "github.com/Liviu2018/employee/EmployeeManagement/api"
	"github.com/Liviu2018/employee/EmployeeManagement/database"
)

func main() {
	database.Init()
	defer database.Close()

	r := mux.NewRouter().StrictSlash(true)

	r.Path("/createEmployee").Methods("POST").HandlerFunc(resthandlers.CreateEmployee)
	r.Path("/listAllEmployees").Methods("GET").HandlerFunc(resthandlers.ListAllEmployees)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	http.Handle("/", r)

	go http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)

	// this way main waits forever, giving a chance to its goroutine to serve incoming API calls; afterwards it can still print a message
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig

	fmt.Println("Receive interrupt signal")
}
