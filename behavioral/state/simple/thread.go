package simple

import "fmt"

// ThreadContext 线程上下文
type ThreadContext struct {
	State State
	newState *NewState
	runnableState *RunnableState
	runningState *RunningState
	blockState *BlockState
	deadState *DeadState
}

func (t *ThreadContext) Start() {
	if t.State.GetState()=="new" {
		t.State.Start()
		t.State = t.runnableState
	}else {
		fmt.Println("unsupported operation")
	}
	fmt.Println(">>>",t.State.GetState())
}

func (t *ThreadContext) Resume() {
	if t.State.GetState()=="block" {
		t.State.Resume()
		t.State = t.runnableState
	}else {
		fmt.Println("unsupported operation")
	}
	fmt.Println(">>>",t.State.GetState())
}

func (t *ThreadContext) Suspend() {
	if t.State.GetState()=="running" {
		t.State.Suspend()
		t.State = t.blockState
	}else {
		fmt.Println("unsupported operation")
	}
	fmt.Println(">>>",t.State.GetState())
}

func (t *ThreadContext) Stop() {
	if t.State.GetState()=="running" {
		t.State.Stop()
		t.State = t.deadState
	}else {
		fmt.Println("unsupported operation")
	}
	fmt.Println(">>>",t.State.GetState())
}

func (t *ThreadContext) GetTimeSlice() {
	if t.State.GetState() == "runnable" {
		t.State.GetTimeSlice()
		t.State = t.runningState
	}else {
		fmt.Println("unsupported operation")
	}
	fmt.Println(">>>",t.State.GetState())
}

func NewThreadContext() *ThreadContext {
	return &ThreadContext{State: &NewState{}}
}

