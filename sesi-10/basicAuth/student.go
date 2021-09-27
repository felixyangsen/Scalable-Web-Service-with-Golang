package main

import "net/http"

var students = []*Student{}

var teachers = []*Teacher{}

type Student struct {
	ID    string
	Name  string
	Grade int
}

type Teacher struct {
	ID   string
	Name string
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

func GetTeachers() []*Teacher {
	return teachers
}

func SelectTeacher(id string) *Teacher {
	for _, each := range teachers {
		if each.ID == id {
			return each
		}
	}

	return nil
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		outputJSON(w, SelectStudent(id))
		return
	}

	outputJSON(w, GetStudents())
}

func ActionTeacher(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		outputJSON(w, SelectTeacher(id))
		return
	}

	outputJSON(w, GetTeachers())
}

func init() {
	students = append(students, &Student{ID: "S001", Name: "Ethan", Grade: 2})
	students = append(students, &Student{ID: "S002", Name: "Bourge", Grade: 2})
	students = append(students, &Student{ID: "S003", Name: "Wick", Grade: 3})

	teachers = append(teachers, &Teacher{ID: "T001", Name: "John"})
	teachers = append(teachers, &Teacher{ID: "T002", Name: "David"})
	teachers = append(teachers, &Teacher{ID: "T003", Name: "Ellen"})
}
