/**
	------------Answered Questions------------

	Q: 	Are there any indexes that would help with the below tasks?
	A: 	Yes, we should have primary key indexes so that fetching the rows when
		when querying the table is more efficient

	Q:	What tables do each database_1 and database_2 have? Explain your reasoning.
	A:	These databases both have student and enrollment tables, which are named the same
		to allow for dynamic querying by swapping out the database name depending on the student(s)
		parameter that is/are currently being queried against. This is a database architectural
		concept known as "sharding". The courses table exists in a separate db (database_3),
		so that it's not tied to either database_1 or database_2 since it is used by both.

 */


package main

import (
	"entities"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println(add(1,2))
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

// helper method for identifying which db a student can/should be located in
func identifyDb(studentName string) string {

	// 0-12th idx == A-M

	var dbName string

	var alpha = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if indexOf(string(studentName[0]), alpha) <= 12 {
		dbName = "[database_1]"
	} else {
		dbName = "[database_2]"
	}

	return dbName
}

func getCoursesByStudent(studentName string) ([]entities.Course, error) {
	// 0-12th idx == A-M

	dbName := identifyDb(studentName)

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

func enrollStudent(name string, number string, courses []entities.Course) error {

	dbName := identifyDb(name)

	var courseIDs string

	for idx := range courses {
		if idx > 0 {
			courseIDs += "," + strconv.FormatInt(courses[idx].CourseID, 10)

		} else {
			courseIDs += strconv.FormatInt(courses[idx].CourseID, 10)
		}
	}

	query := `
		SET NOCOUNT ON
		SET XACT_ABORT ON

		BEGIN TRY
			BEGIN TRAN

			INSERT INTO ` + dbName + `.[dbo].[Students] (
				[StudentName],
				[StudentPhoneNumber]
			)
			VALUES (
				` + name + `,
				` + number + `
			)

			DECLARE @StudentID INT = SCOPE_IDENTITY()
			
			INSERT INTO ` + dbName + `.[dbo].[Enrollment] (
				[Enrollment_StudentID],
				[Enrollment_CourseID],
				[DateEnrolled],
				[FinalGrade]
			)
			SELECT
				@StudentID,
				[value] AS [CourseID],
				GETDATE() AS [DateEnrolled]',
				NULL AS [FinalGrade]
			FROM STRING_SPLIT(` + courseIDs + `, ',')

			COMMIT TRAN
		END TRY
		
		BEGIN CATCH
		
			IF XACT_STATE() = -1
			BEGIN
				ROLLBACK TRAN
			END
			IF XACT_STATE() = 1
			BEGIN
				COMMIT TRAN
			END

		END CATCH
	`

	err := executeEnrollStudentSql(dbName, query)
	if err != nil {
		return err
	}

	return nil
}


// bogus executor to resolve errors in above function
func executeEnrollStudentSql(dbName string, query string) error {
	return nil
}

func getCoursesForStudents(students []entities.Student) map[entities.Student][]entities.Course {

	// initialize return shape
	studentCourseMap := make(map[entities.Student][]entities.Course)

	// loop through students
	for idx := range students {

		student := students[idx]

		dbName := identifyDb(student.StudentName)

		query := `
			SELECT 
				[CourseID],
				[CourseCode],
				[CourseName]
			FROM [database_3].[dbo].[Courses] c
			INNER JOIN  ` + dbName + `.[dbo].[Enrollment] e ON e.[Enrollment_CourseID] = c.[CourseID]
			INNER JOIN ` + dbName + `.[dbo].[Students] s ON s.[StudentID] = e.[Enrollment_StudentID]
			WHERE s.[StudentID] = ` + strconv.FormatInt(student.StudentID, 10)

		coursesForStudent := executeGetCoursesForStudentSql(dbName, query)

		studentCourseMap[student] = coursesForStudent
	}

	return studentCourseMap
}


// bogus helper executor method for above function
// note this only gets the courses for a single student, and the mapping takes place
// in the parent function that calls this one
func executeGetCoursesForStudentSql(dbName string, query string) []entities.Course {

	var courses []entities.Course

	course1 := entities.Course{
		CourseID:   1,
		CourseCode: "00519",
		CourseName: "Mathematics",
	}

	course2 := entities.Course{
		CourseID:   2,
		CourseCode: "00847",
		CourseName: "Chemistry",
	}

	return append(courses, course1, course2)
}