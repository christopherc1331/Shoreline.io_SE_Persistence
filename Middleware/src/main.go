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

// example input: "ALGORITHMS1"
func getStudentsInCourse(courseCode string) ([]entities.Student, error) {

	var students []entities.Student

	var query1 string
	query1 = `
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

	var query2 string
	query2 = `
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