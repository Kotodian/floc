package floc

type Job func(ctx Context, ctrl Control) error
