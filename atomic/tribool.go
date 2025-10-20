package atomic

import "sync/atomic"

// AtomicTriBool is an atomic tri-state boolean value that can be
// true, false, or unset.
//
// The zero value is unset.
type AtomicTriBool struct {
	// 0: unset, 1: true, 2: false
	value atomic.Uint64
}

func (atb *AtomicTriBool) Get() (v bool, set bool) {
	val := atb.value.Load()
	switch val {
	case 1:
		return true, true
	case 2:
		return false, true
	}

	// unset
	return false, false
}

func (atb *AtomicTriBool) Set(v bool) {
	var newVal uint64
	if v {
		newVal = 1
	} else {
		newVal = 2
	}

	atb.value.Store(newVal)
}
