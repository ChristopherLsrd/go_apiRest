package rest

import(
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"internal/entities"
	"internal/persistence"
	"encoding/json"
	"log"
	"strconv"
	"io/ioutil"
)

var studentDAOMem=persistence.NewStudentDAOMem()

//students/{id}
func StudentHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])

//	id, ok := []

	data := studentDAOMem.Find(id)

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println("An error occurred while marshaling weather in json format: ", err)
		return
	}

	fmt.Fprintf(w, "%s", j)

}

//students
func StudentsHandler(w http.ResponseWriter, r *http.Request){
		
		j, err := json.Marshal(studentDAOMem.FindAll())
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s \n", j)	

	
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var student entities.Student
    json.Unmarshal(reqBody, &student)
	studentDAOMem.Create(student)
 
    json.NewEncoder(w).Encode(student)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id,err := strconv.Atoi(vars["id"])
	if err!=nil{
		log.Fatal(err)
	}
 
   studentDAOMem.Delete(id)
 
}

func PutStudentHandler(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
	
    var student entities.Student
	
    json.Unmarshal(reqBody, &student)
	studentDAOMem.Update(student)
}
