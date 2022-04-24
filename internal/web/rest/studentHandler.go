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

// swagger:operation GET /students/{id} Students GetStudent
// ---
// summary: Return student.
// description: Return a student if id provided match with one student in the database.
// parameters:
// - name: id
//   in: path
//   description: student id
//   type: string
//   required: true
// responses:
//   "200": Student returned

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

// swagger:operation GET /students Students GetStudents
// ---
// summary: Return students.
// description: Return all the studens.
// responses:
//   "200": Students returned

func StudentsHandler(w http.ResponseWriter, r *http.Request) {

	//j, err := json.Marshal(studentDAOMem.FindAll())
	j, err := json.Marshal(studentDAOBolt.FindAll())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "%s \n", j)

}

// swagger:operation POST /students Students PostStudent
// ---
// summary: Add new student.
// description: Add a new student if id provided not existing.
// parameters:
// - name: student
//   description: student to add int the student bucket
//   in: body
//   required: true
//   schema:
//     "$ref": "#/internal/entities/Student"
// responses:
//   "200": Student added
func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var student entities.Student
	json.Unmarshal(reqBody, &student)
	//studentDAOMem.Create(student)
	studentDAOBolt.Create(student)
	json.NewEncoder(w).Encode(student)
}

// swagger:operation DELETE /students/{id} Students DeleteStudent
// ---
// summary: Delete a student.
// description: Delete a student if id provided existing.
// parameters:
// - name: id
//   description: student id
//   in: path
//   required: true
// responses:
//   "200": Student deleted
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	//studentDAOMem.Delete(id)
	studentDAOBolt.Delete(id)

}

// swagger:operation PUT /students Students PutStudent
// ---
// summary: Modify an existing student.
// description: Modify an existing student.
// parameters:
// - name: student
//   description: student to modify
//   in: body
//   required: true
//   schema:
//     "$ref": "#/internal/entities/Student"
// responses:
//   "200": Student modified

func PutStudentHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var student entities.Student

	json.Unmarshal(reqBody, &student)
	//studentDAOMem.Update(student)
	studentDAOBolt.Update(student)
}
