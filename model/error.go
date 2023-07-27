package model

import "fmt"

type Error struct {
	Code       int
	Err        error
	Who        string
	StatusHTTP int
	Data       interface{}
	APIMessage string
	UserID     string
}

func (e *Error) Error() string {
	return fmt.Sprintf(
		"Code: %d, Who: %s, StatusHTTP: %d, Data: %v, APIMessage: %s, UserID: %s, Err: %v",
	)
}
