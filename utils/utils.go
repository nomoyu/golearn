package utils

import "fmt"

func AddError(existError, newError error) error {
	if existError == nil {
		return newError
	}
	return fmt.Errorf("%v,%w", existError, newError)
}
