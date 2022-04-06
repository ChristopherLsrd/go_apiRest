package rest

import(
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"internal/entities"
	"encoding/json"
	"log"
	"strconv"
	"io/ioutil"
)

var students []entities.Student = []entities.Student{
    entities.NewStudent(1, "Gaspar", "Missiaen", 21, "23"), entities.NewStudent(2, "Daurian", "Jsp", 20, "Go"), entities.NewStudent(3, "Daryl", "Jsp", 20, "-2"), entities.NewStudent(4, "Christopher", "Lessirard", 20, "26"),
}

//students/{id}
func StudentHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])

//	id, ok := []

	data := getStudent(id)

	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println("An error occurred while marshaling weather in json format: ", err)
		return
	}

	fmt.Fprintf(w, "%s", j)

}

//students
func StudentsHandler(w http.ResponseWriter, r *http.Request){
		
		students=append(students)
		j, err := json.Marshal(students)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s \n", j)	

	
}

func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
    var student entities.Student
    json.Unmarshal(reqBody, &student)
    students = append(students, student)
 
    json.NewEncoder(w).Encode(student)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id,err := strconv.Atoi(vars["id"])
	if err!=nil{
		log.Fatal(err)
	}
 
    for index, student := range students {
        if student.Id == id {
            students = append(students[:index], students[index+1:]...)
        }
    }
 
}

func PutStudentHandler(w http.ResponseWriter, r *http.Request) {
    reqBody, _ := ioutil.ReadAll(r.Body)
	
    var student entities.Student
	
    json.Unmarshal(reqBody, &student)
	for i,studentFor:=range students{
		if student.Id==studentFor.Id{
			students[i]=student
		}
	}
}

func getStudent(id int) entities.Student{
	for _,student:= range students{
		if student.Id == id{
			return student
		}
	}
	
	return entities.NewStudent(0,"","",0,"")
}