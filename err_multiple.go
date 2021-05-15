package floc

import (
	"strconv"
	"strings"
)

type ErrMultiple struct {
	list []error
}

func NewErrMultiple(errs ...error) ErrMultiple {
	return ErrMultiple{list: errs}
}

func (err ErrMultiple) Top() error {
	if len(err.list) > 0 {
		return err.list[0]
	}
	return nil
}

func (err ErrMultiple) List() []error {
	return err.list
}
func (err ErrMultiple) Len() int {
	return len(err.list)
}

func (err ErrMultiple) Error() string {
	if len(err.list) == 1 {
		return err.list[0].Error()
	}
	sb := strings.Builder{}
	sb.WriteString(strconv.Itoa(len(err.list)))
	sb.WriteString(" errors: ")
	for i, err := range err.list {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteByte('"')
		sb.WriteString(err.Error())
		sb.WriteByte('"')
	}
	return sb.String()
}
