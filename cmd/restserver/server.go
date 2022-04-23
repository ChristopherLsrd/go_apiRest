package main

import (
	"internal/web/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	//DAOMem
	r.HandleFunc("/rest/languages/{code}", rest.LanguageHandler).Methods("GET")
	r.HandleFunc("/rest/languages/{code}", rest.DeleteLanguageHandler).Methods("DELETE")
	r.HandleFunc("/rest/languages", rest.LanguagesHandler).Methods("GET")
	r.HandleFunc("/rest/languages", rest.PostLanguageHandler).Methods("POST")
	r.HandleFunc("/rest/languages", rest.PutLanguageHandler).Methods("PUT")
	r.HandleFunc("/rest/students/{id}", rest.StudentHandler).Methods("GET")
	r.HandleFunc("/rest/students/{id}", rest.DeleteStudentHandler).Methods("DELETE")
	r.HandleFunc("/rest/students", rest.StudentsHandler).Methods("GET")
	r.HandleFunc("/rest/students", rest.PostStudentHandler).Methods("POST")
	r.HandleFunc("/rest/students", rest.PutStudentHandler).Methods("PUT")

	//DAOBolt

	log.Fatal(http.ListenAndServe(":8080", r))

}
