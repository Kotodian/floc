package floc

type ErrInvalidJob struct{}

const invalidJobMessage = "job is invalid"

func (ErrInvalidJob) Error() string {
	return invalidJobMessage
}
