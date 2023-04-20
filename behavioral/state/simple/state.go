package simple

import "fmt"

// NewState 新建
type NewState struct {
	State
}

func (state *NewState) Start() {
	fmt.Println("start...")
}

func (state *NewState) GetState() string{
	return "new"
}

// RunnableState 就绪
type RunnableState struct {
	State
}

func (state *RunnableState) GetState() string{
	return "runnable"
}

func (state *RunnableState) GetTimeSlice() {
	fmt.Println("get time slice...")
}

// RunningState 运行中
type RunningState struct {
	State
}
func (state *RunningState) GetState() string{
	return "running"
}
func (state *RunningState) Suspend()  {
	fmt.Println("suspend...")
}

func (state *RunningState) Stop()  {
	fmt.Println("stop...")
}

// BlockState 阻塞
type BlockState struct {
	State
}
func (state *BlockState) GetState() string{
	return "block"
}
func (state *BlockState) Resume()  {
	fmt.Println("resume...")
}

// DeadState 死亡
type DeadState struct {
	State
}
func (state *DeadState) GetState() string{
	return "dead"
}

