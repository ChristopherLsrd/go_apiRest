package main

import (
	/*"internal/entities"
	"internal/persistence/bolt"*/
	"internal/web/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	/*	var languages []entities.Language = []entities.Language{
			entities.NewLanguage("21", "Go"), entities.NewLanguage("12", "Python"),
		}
		var b = bolt.DBopen("base.db")
		for _, l := range languages {
			b.DBput("languages", l.Code, l.String())
		}*/

	r := mux.NewRouter()

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

	log.Fatal(http.ListenAndServe(":8080", r))

}
