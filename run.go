package floc

func Run(job Job) (result Result, data interface{}, err error) {
	ctx := NewContext()
	defer ctx.Release()

	ctrl := NewControl(ctx)
	defer ctrl.Release()

	return RunWith(ctx, ctrl, job)
}

func RunWith(ctx Context, control Control, job Job) (result Result, data interface{}, err error) {
	if job == nil {
		return None, nil, ErrInvalidJob{}
	}

	unhandledErr := job(ctx, control)
	result, data, err = control.Result()
	if result != None {
		return result, data, err
	}

	if unhandledErr != nil {
		return Failed, nil, unhandledErr
	}
	return None, nil, nil
}
