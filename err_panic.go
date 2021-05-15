package floc

import "fmt"

const errorPrefix = "panic with"

type ErrPanic struct {
	data interface{}
}

func NewErrPanic(data interface{}) ErrPanic {
	return ErrPanic{data}
}

func (err ErrPanic) Data() interface{} {
	return err.data
}

func (err ErrPanic) Error() string {
	switch v := err.data.(type) {
	case error:
		return errorPrefix + v.Error()
	case fmt.Stringer:
		return errorPrefix + v.String()
	default:
		return fmt.Sprintf("%s%v", errorPrefix, err.data)
	}
}
