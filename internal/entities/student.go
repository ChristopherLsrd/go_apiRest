package entities

import (
	"fmt"
)

type Student struct {
	Id           int
	FirstName    string
	LastName     string
	Age          int
	LanguageCode string
}

func NewStudent(id int, fname string, lname string, age int, languageCode string) Student {
	student := Student{
		Id:           id,
		FirstName:    fname,
		LastName:     lname,
		Age:          age,
		LanguageCode: languageCode,
	}
	return student
}

func (s Student) String() string {
	return fmt.Sprintf("%d %s %s %d %s", s.Id, s.FirstName, s.LastName, s.Age, s.LanguageCode)
}
