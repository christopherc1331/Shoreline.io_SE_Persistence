USE [master]
GO

IF DB_ID('database_1') IS NULL
BEGIN
	CREATE DATABASE [database_1]
END

IF DB_ID('database_2') IS NULL
BEGIN
	CREATE DATABASE [database_2]
END