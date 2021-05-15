package floc

import (
	"bytes"
	"fmt"
)

type ResultMask Result

const emptyResultMask = ResultMask(0)

func NewResultMask(mask Result) ResultMask {
	return ResultMask(mask & usedBitsMask)
}

func EmptyResultMask() ResultMask {
	return emptyResultMask
}

func (mask ResultMask) IsMasked(result Result) bool {
	return mask&ResultMask(result) == ResultMask(result)
}

func (mask ResultMask) IsEmpty() bool {
	return mask == 0
}

func (mask ResultMask) String() string {
	buf := &bytes.Buffer{}

	_, _ = fmt.Fprint(buf, "[")
	empty := true
	for _, result := range []Result{None, Completed, Canceled, Failed} {
		if mask.IsMasked(result) {
			if empty {
				fmt.Fprint(buf, result.String())
			} else {
				fmt.Fprint(buf, ",", result.String())
			}
			empty = false
		}
	}
	fmt.Fprint(buf, "]")
	return buf.String()
}
