package lwha

import "github.com/satori/go.uuid"

func Uuid()string{
	v4 := uuid.NewV4()
	s := v4.String()
	return s
}

