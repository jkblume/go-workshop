package foo

import (
	"0020-handle-errors-gracefully/internal/database"
	"fmt"
)

func GetExample(throwError bool) (string, error) {
	value1, err := database.ReadValue("1", throwError)
	if err != nil {
		return "", err
	}

	value2, err := database.ReadValue("2", throwError)
	if err != nil {
		return "", err
	}

	return value1 + "/" + value2, err
}

func GetExample2(throwError bool) (string, error) {
	value1, err := database.ReadValue("1", throwError)
	if err != nil {
		return "", fmt.Errorf("failed to read value 1: %v", err)
	}

	value2, err := database.ReadValue("2", throwError)
	if err != nil {
		return "", fmt.Errorf("failed to read value 2: %v", err)
	}

	return value1 + "/" + value2, err
}

func GetExample3(throwError bool) (string, error) {
	value1, err := database.ReadValue("1", throwError)
	if err != nil {
		return "", fmt.Errorf("failed to read value 1: %w", err)
	}

	value2, err := database.ReadValue("2", throwError)
	if err != nil {
		return "", fmt.Errorf("failed to read value 2: %w", err)
	}

	return value1 + "/" + value2, err
}
