package rest

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/persistence"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var studentDAOMem = persistence.NewStudentDAOMem()
var studentDAOBolt = persistence.NewStudentDAOBolt()

//students/{id}
func StudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	//data := studentDAOMem.Find(id)
	data := studentDAOBolt.Find(id)

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "%s", j)

}

//students
func StudentsHandler(w http.ResponseWriter, r *http.Request) {

	//j, err := json.Marshal(studentDAOMem.FindAll())
	j, err := json.Marshal(studentDAOBolt.FindAll())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s \n", j)

}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var student entities.Student
	json.Unmarshal(reqBody, &student)
	//studentDAOMem.Create(student)
	studentDAOBolt.Create(student)
	json.NewEncoder(w).Encode(student)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	//studentDAOMem.Delete(id)
	studentDAOBolt.Delete(id)

}

func PutStudentHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var student entities.Student

	json.Unmarshal(reqBody, &student)
	//studentDAOMem.Update(student)
	studentDAOBolt.Update(student)
}
