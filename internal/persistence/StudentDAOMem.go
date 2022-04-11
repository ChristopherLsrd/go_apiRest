package persistence

import(
	"internal/entities"
	"sort"
)

var students []entities.Student = []entities.Student{
    entities.NewStudent(2, "Daurian", "Jsp", 20, "Go"), entities.NewStudent(3, "Daryl", "Jsp", 20, "-2"), entities.NewStudent(4, "Christopher", "Lessirard", 20, "26"),entities.NewStudent(1, "Gaspar", "Missiaen", 21, "23"),
}

type StudentDAOMem struct{
	
}
func NewStudentDAOMem() StudentDAOMem{

	return StudentDAOMem{}
}

func(StudentDAOMem) FindAll() [] entities.Student{
	sort.SliceStable(students, func(i, j int) bool {
		return students[i].Id < students[j].Id
	})
	return students
}

func(StudentDAOMem) Find(Id int) entities.Student{
	for i,studentFor:=range students{
		if studentFor.Id==Id{
			return students[i]
		}
	}
	return entities.Student{0,"","",0,""}
	
}

func(StudentDAOMem) Exists(Id int) bool{
	for _,studentFor:=range students{
		if studentFor.Id==Id{
			return true
		}
	}
	return false
}

func(StudentDAOMem) Delete(Id int) bool{

		for i,studentFor:=range students{
			if studentFor.Id==Id{
				students = append(students[:i], students[i+1:]...)
				return true
			}
		}
		return false
	
}

func(StudentDAOMem) Create(student entities.Student) bool{
	for _,studentFor:=range students{
		if studentFor.Id==student.Id{	
			return false
		}
	}
	students = append(students, student)
	return true
}
	
func(StudentDAOMem) Update(student entities.Student) bool{
	for i,studentFor:=range students{
		if studentFor.Id==student.Id{	
			students[i]=student
			return true
		}
	}
	
	return false
}