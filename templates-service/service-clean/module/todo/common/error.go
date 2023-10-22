package common

import "fmt"

var (
	ErrInvalidStatus = "invalid status"
)

func WrapError(errContext string, err error) error {
	if err == nil {
		return err
	}
	return fmt.Errorf("%s: %w", errContext, err)
}
