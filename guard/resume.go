package guard

import "github.com/Kotodian/floc"

func Resume(mask floc.ResultMask, job floc.Job) floc.Job {
	if mask.IsEmpty() {
		return func(ctx floc.Context, ctrl floc.Control) error {
			mockCtx := floc.NewContext()
			defer mockCtx.Release()

			mockCtrl := floc.NewControl(mockCtx)
			defer mockCtrl.Release()

			return job(mockContext{ctx, mockCtx}, mockCtrl)
		}
	}

	return func(ctx floc.Context, ctrl floc.Control) error {
		mockCtx := mockContext{Context: ctx, mock: floc.NewContext()}
		mockCtrl := floc.NewControl(mockCtx)
		defer func() {
			mockCtrl.Release()
			mockCtx.Release()

			if mockCtrl.IsFinished() {
				res, data, err := mockCtrl.Result()
				if !mask.IsMasked(res) {
					switch res {
					case floc.Canceled:
						ctrl.Cancel(data)
					case floc.Completed:
						ctrl.Complete(data)
					case floc.Failed:
						ctrl.Fail(data, err)
					}
				}
			}
		}()
		return job(mockCtx, mockCtrl)
	}
}
