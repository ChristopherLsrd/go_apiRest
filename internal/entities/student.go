package entities

import(
	"fmt"
)

type Student struct{
	Id int
	FirstName string
	LastName string
	Age int
	LanguageCode string
}

func newStudent(fname string,lname string,age int,languageCode string) Student{
	student:=Student{
		Id:1,
		FirstName:fname,
		LastName:lname,
		Age:age,
		LanguageCode:languageCode,
	}
	return student
}


func (s Student) String() string{
	return fmt.Sprintf("%s %s",s.FirstName, s.LastName )
}