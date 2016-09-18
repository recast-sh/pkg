package core

import "fmt"

var (
	ErrExpectedValue         = Error("Expected Value")
	ErrMissingAuthentication = Error("Missing authentication credentials")

	ErrMustUsePlan         = Error("Must use Plan")
	ErrPlanNested          = Error("Plan() cannot be nested")
	ErrPlanAlreadyExecuted = Error("Plan already Executed")

	ErrNotInUnitsScope = Error("Not in Units() scope")
	ErrUnitsNested     = Error("Units() cannot be nested")

	ErrNotInFilesScope = Error("Not in Files() scope")
	ErrFilesNested     = Error("Files() cannot be nested")

	ErrNotInPasswdScope = Error("Not in Passwd() scope")
	ErrPasswdNested     = Error("Passwd() cannot be nested")
)

type RecastError interface {
	error
	// No-op but marks error for better handling
	RecastError()
}

func IsRecastError(err interface{}) bool {
	_, ok := err.(RecastError)
	return ok
}

type recastError struct {
	message string
}

func (e recastError) Error() string {
	return e.message
}

func (e recastError) RecastError() {
}

func Error(msg string) RecastError {
	return recastError{
		message: msg,
	}
}

func Errorf(f string, args ...interface{}) RecastError {
	return recastError{
		message: fmt.Sprintf(f, args...),
	}
}
