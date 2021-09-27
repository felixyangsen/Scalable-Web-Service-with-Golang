package main

var students = []*Student{}

type Student struct {
	ID    string
	name  string
	Grade int
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.ID == id {
			return each
		}
	}

	return nil
}

func init() {
	students = append(students, &Student{ID: "S001", name: "Ethan", Grade: 2})
	students = append(students, &Student{ID: "S002", name: "Bourge", Grade: 2})
	students = append(students, &Student{ID: "S003", name: "Wick", Grade: 3})
}
