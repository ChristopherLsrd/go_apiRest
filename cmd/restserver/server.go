// API of students and languages.
//
// Terms Of Service:
//     Schemes: http, https
//     Host: localhost:8080
//     Version: 1.0.0
//	   BasePath: /rest
//     Contact: Christopher Lessirard <christopher.lessirard@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"internal/web/rest"
	"log"
	"net/http"
)

func main() {

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

	r.Handle("/swagger.json", http.FileServer(http.Dir("../../docSwagger")))

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.json"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)

	log.Fatal(http.ListenAndServe(":8080", r))

}
