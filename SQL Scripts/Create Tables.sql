USE [database_3]
GO

-- ======== Courses ========

IF OBJECT_ID(N'[database_3].[dbo].[Courses]', N'U') IS NOT NULL
BEGIN

	CREATE TABLE [database_3].[dbo].[Courses] (
		[CourseID] INT IDENTITY(1,1) PRIMARY KEY,
		[CourseCode] VARCHAR(50) NOT NULL,
		[CourseName] VARCHAR(400) NOT NULL
	)

END

GO


-- ======== Student-Related (A-M) ========

USE [database_1]
GO

IF OBJECT_ID(N'[database_1].[dbo].[Students]', N'U') IS NOT NULL
BEGIN

	CREATE TABLE [database_1].[dbo].[Students] (
		[StudentID] INT IDENTITY(1,1) PRIMARY KEY,
		[StudentName] VARCHAR(200) NOT NULL,
		[StudentPhoneNumber] VARCHAR(50) NOT NULL  -- varchar for now because I'm not sure if hyphens/parenthesis/spaces are allowed
	)

END


IF OBJECT_ID(N'[database_1].[dbo].[Enrollment]', N'U') IS NOT NULL
BEGIN

	CREATE TABLE [database_1].[dbo].[Enrollment] (
		[EnrollmentID] INT IDENTITY(1,1) PRIMARY KEY,
		[Enrollment_StudentID] INT NOT NULL FOREIGN KEY REFERENCES [database_1].[dbo].[Students] ([StudentID]),
		[Enrollment_CourseID] INT NOT NULL FOREIGN KEY REFERENCES [database_3].[dbo].[Courses] ([CourseID]),
		[DateEnrolled] DATE NOT NULL,
		[FinalGrade] FLOAT NULL -- allowing for null in the case that someone may want to enter info first, 
								 --	then update the final grade at a later time
	)

END

GO


-- ======== Student-Related (N-Z) ========

USE [database_2]
GO

IF OBJECT_ID(N'[database_2].[dbo].[Students]', N'U') IS NOT NULL
BEGIN

	CREATE TABLE [database_2].[dbo].[Students] (
		[StudentID] INT IDENTITY(1,1) PRIMARY KEY,
		[StudentName] VARCHAR(200) NOT NULL,
		[StudentPhoneNumber] VARCHAR(50) NOT NULL  -- varchar for now because I'm not sure if hyphens/parenthesis/spaces are allowed
	)

END


IF OBJECT_ID(N'[database_2].[dbo].[Enrollment]', N'U') IS NOT NULL
BEGIN

	CREATE TABLE [database_2].[dbo].[Enrollment] (
		[EnrollmentID] INT IDENTITY(1,1) PRIMARY KEY,
		[Enrollment_StudentID] INT NOT NULL FOREIGN KEY REFERENCES [database_2].[dbo].[Students] ([StudentID]),
		[Enrollment_CourseID] INT NOT NULL FOREIGN KEY REFERENCES [database_3].[dbo].[Courses] ([CourseID]),
		[DateEnrolled] DATE NOT NULL,
		[FinalGrade] FLOAT NULL -- allowing for null in the case that someone may want to enter info first, 
								 --	then update the final grade at a later time
	)

END

GO
