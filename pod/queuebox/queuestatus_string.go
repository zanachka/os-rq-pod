// Code generated by "stringer -type=QueueStatus -linecomment"; DO NOT EDIT.

package queuebox

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Working-1]
	_ = x[Paused-2]
}

const _QueueStatus_name = "workingpaused"

var _QueueStatus_index = [...]uint8{0, 7, 13}

func (i QueueStatus) String() string {
	i -= 1
	if i < 0 || i >= QueueStatus(len(_QueueStatus_index)-1) {
		return "QueueStatus(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _QueueStatus_name[_QueueStatus_index[i]:_QueueStatus_index[i+1]]
}
