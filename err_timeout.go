package floc

import (
	"fmt"
	"time"
)

type ErrTimeout struct {
	id interface{}
	at time.Time
}

const tplTimeoutMessage = "%v timed out at %s"

func NewErrTimeout(id interface{}, at time.Time) ErrTimeout {
	return ErrTimeout{id, at}
}

func (err ErrTimeout) ID() interface{} {
	return err.id
}

func (err ErrTimeout) At() time.Time {
	return err.at
}

func (err ErrTimeout) Error() string {
	return fmt.Sprintf(tplTimeoutMessage, err.id, err.at)
}
