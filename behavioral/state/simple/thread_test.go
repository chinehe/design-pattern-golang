package simple

import (
	"testing"
)

func TestThreadState(t *testing.T) {
	thread := NewThreadContext()
	thread.Start()
	thread.GetTimeSlice()
	thread.GetTimeSlice()
	thread.Suspend()
	thread.Resume()
	thread.GetTimeSlice()
	thread.Stop()
}
