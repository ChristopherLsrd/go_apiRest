package persistence

import (
	"encoding/json"
	"fmt"
	"internal/entities"
	"internal/persistence/bolt"
	"log"
)

var bS = bolt.GetboltDB()

type StudentDAOBolt struct{}

func NewStudentDAOBolt() StudentDAOBolt {
	return StudentDAOBolt{}
}

func (s StudentDAOBolt) FindAll() []entities.Student {
	res := bS.DBgetAll("Students")
	var students []entities.Student
	for _, l := range res {

		var student entities.Student
		json.Unmarshal([]byte(l), &student)
		students = append(students, student)
	}
	return students
}

func (s StudentDAOBolt) Find(id int) entities.Student {
	var student entities.Student
	res := bS.DBget("Students", fmt.Sprintf("%d", id))

	if res == "" {
		fmt.Println("Id :", id, "non trouvé")
	}

	json.Unmarshal([]byte(res), &student)
	return student

}

func (s StudentDAOBolt) Delete(id int) bool {

	err := bS.DBdelete("Students", fmt.Sprintf("%d", id))
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (s StudentDAOBolt) Create(student entities.Student) bool {
	res, _ := json.Marshal(student)
	id := fmt.Sprintf("%d", student.Id)
	if s.Exists(student.Id) == false {
		bS.DBput("Students", id, string(res))
		return true
	}
	return false

}

func (s StudentDAOBolt) Exists(id int) bool {
	student := bS.DBget("Students", fmt.Sprintf("%d", id))
	if student != "" {
		return true
	}

	return false
}

func (s StudentDAOBolt) Update(student entities.Student) bool {
	res, _ := json.Marshal(student)
	id := fmt.Sprintf("%d", student.Id)
	if s.Exists(student.Id) {
		bS.DBdelete("Students", id)
		bS.DBput("Students", id, string(res))
		return true
	}
	return false

}
