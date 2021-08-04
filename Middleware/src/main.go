package main

import (
	"entities"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println(add(1,2))
	fmt.Println(getStudentsInCourse("ALGORITHMS1"))
}

// this is a sanity check function for my own self teaching,
// since I've never worked with GO prior to this

func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}


// =========== actual solutions below ===========

// helper method for identifying index of char in array
func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}

func getCoursesByStudent(studentName string) ([]entities.Course, error) {
	// 0-12th idx == A-M

	var dbName string

	var alpha = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if indexOf(string(studentName[0]), alpha) <= 12 {
		dbName = "[database_1]"
	} else {
		dbName = "[database_2]"
	}
	query := `
		SELECT 
			[CourseID],
			[CourseCode],
			[CourseName]
		FROM [database_3].[dbo].[Courses] c
		INNER JOIN  ` + dbName + `.[dbo].[Enrollment] e ON e.[Enrollment_CourseID] = c.[CourseID]
		INNER JOIN ` + dbName + `.[dbo].[Students] s ON s.[StudentID] = e.[Enrollment_StudentID]
		WHERE s.[StudentName] = ` + studentName

	return executeGetCoursesByStudentSql(dbName, query), nil
}


// bogus executor to resolve errors in above function
func executeGetCoursesByStudentSql(dbName string, query string) []entities.Course {

	var courses []entities.Course

	course1 := entities.Course{
		CourseID:   1,
		CourseCode: "00519",
		CourseName: "Mathematics",
	}

	return append(courses, course1)
}

// example input: "ALGORITHMS1"
func getStudentsInCourse(courseCode string) ([]entities.Student, error) {

	var students []entities.Student

	query1 := `
		SELECT 
			s.[StudentID],
			s.[StudentName],
			s.[StudentPhoneNumber]
		FROM [database_3].[dbo].[Courses] c
		INNER JOIN [database_1].[dbo].[Enrollment] e ON e.[Enrollment_CourseID] = c.[CourseID]
		INNER JOIN [database_1].[dbo].[Students] s ON s.[StudentID] = e.[Enrollment_StudentID]
		WHERE c.[CourseCode] = ` + courseCode

	query1Results := executeGetStudentsInCourseSql("[database_1]", query1)

	students = append(students, query1Results...)

	query2 := `
		SELECT 
			s.[StudentID],
			s.[StudentName],
			s.[StudentPhoneNumber]
		FROM [database_3].[dbo].[Courses] c
		INNER JOIN [database_2].[dbo].[Enrollment] e ON e.[Enrollment_CourseID] = c.[CourseID]
		INNER JOIN [database_2].[dbo].[Students] s ON s.[StudentID] = e.[Enrollment_StudentID]
		WHERE c.[CourseCode] = ` + courseCode

	query2Results := executeGetStudentsInCourseSql("[database_2]", query2)

	students = append(students, query2Results...)

	return students, nil
}


// bogus executor to resolve errors in above function
// is hard coded and will produce duplicate appends to outer function result, but in
// a real life scenario the dbName and query would actually be implemented and persisting
// to db call
func executeGetStudentsInCourseSql(dbName string, query string) []entities.Student {

	var students []entities.Student

	student1 := entities.Student{
		StudentID:          1,
		StudentName:        "Foo",
		StudentPhoneNumber: "1234567890",
	}

	student2 := entities.Student{
		StudentID:          2,
		StudentName:        "Bar",
		StudentPhoneNumber: "0987654321",
	}

	students = append(students, student1, student2)

	return students
}