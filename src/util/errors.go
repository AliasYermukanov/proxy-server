package util

import "fmt"

type argError struct {
	Code             int    `json:"code"`
	DeveloperMessage string `json:"developerMessage"`
}

func (e *argError) Error() string {
	return fmt.Sprintf("%q %s", e.Code, e.DeveloperMessage)
}

var (
	CommonError = &argError{500, "dev"}
	NoFound     = &argError{404, "dev"}
)

func (e *argError) SetDevMessage(developMessage string) *argError {
	e.DeveloperMessage = developMessage
	return e
}
