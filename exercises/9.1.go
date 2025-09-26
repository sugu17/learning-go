package exercises

import "errors"

var (
	ErrInvalidID = errors.New("Invalid ID")
)

func FindItem_9_1(id string) (string, error) {
	if id == "" {
		return "", ErrInvalidID
	}
	return "Never gonna give you up", nil
}
