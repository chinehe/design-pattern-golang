package state

// State 状态接口
type State interface {
	Start()
	Resume()
	Suspend()
	Stop()
	GetTimeSlice()
	GetState() string
}
