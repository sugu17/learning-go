package exercises

import "fmt"

type MissingFieldError struct {
	FieldName string
}

func (me MissingFieldError) Error() string {
	return fmt.Sprintf("Field %v is missing", me.FieldName)
}

type Employee struct {
	Name           string
	JoinedAt       string
	AttendanceConf *map[string]string
}

func SaveEmployeeDescription_9_2(employee Employee) (bool, error) {
	if employee.AttendanceConf == nil {
		return false, MissingFieldError{FieldName: "AttendanceConf"}
	}
	return true, nil
}
