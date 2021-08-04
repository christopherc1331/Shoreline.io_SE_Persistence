package entities

type Enrollment struct {
	EnrollmentID		int64
	EnrollmentStudentID	int64
	EnrollmentCourseID	int64
	DateEnrolled		string
	FinalGrade			float64
}