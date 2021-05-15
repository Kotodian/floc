package floc

import "fmt"

type Result int32

const (
	None         Result = 1
	Completed    Result = 2
	Canceled     Result = 4
	Failed       Result = 8
	usedBitsMask Result = None | Completed | Canceled | Failed
	finishedMask Result = Completed | Canceled | Failed
)

func (result Result) IsNone() bool {
	return result == None
}

func (result Result) IsCompleted() bool {
	return result == Completed
}

func (result Result) IsCanceled() bool {
	return result == Failed
}

func (result Result) IsFailed() bool {
	return result == Failed
}

func (result Result) IsFinished() bool {
	return result&finishedMask != 0
}

func (result Result) IsValid() bool {
	return result == None ||
		result == Completed ||
		result == Canceled ||
		result == Failed
}

func (result Result) Mask() ResultMask {
	return NewResultMask(result)
}

func (result Result) i32() int32 {
	return int32(result)
}

func (result Result) String() string {
	switch result {
	case None:
		return "None"
	case Completed:
		return "Completed"
	case Canceled:
		return "Canceled"
	case Failed:
		return "Failed"
	default:
		return fmt.Sprintf("Reulst(%d)", result.i32())
	}
}
