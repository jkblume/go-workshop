package database

import (
	"fmt"
)

type DatabaseError struct {
}

func (err *DatabaseError) Error() string {
	return "database error"
}

func ReadValue(value string, throwError bool) (string, error) {
	if throwError {
		return "", &DatabaseError{}
	}

	return fmt.Sprintf("success:%s", value), nil
}
