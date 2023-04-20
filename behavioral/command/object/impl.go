package object

// LightOnCommand 开灯命令
type LightOnCommand struct {
	Light *Light
}

func (command *LightOnCommand) Execute() {
	command.Light.On()
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{Light: light}
}

// LightOffCommand 关灯命令
type LightOffCommand struct {
	Light *Light
}

func (command *LightOffCommand) Execute() {
	command.Light.Off()
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{Light: light}
}

// DoorOpenCommand 开门命令
type DoorOpenCommand struct {
	Door *GarageDoor
}

func (command *DoorOpenCommand) Execute() {
	command.Door.Open()
}

func NewDoorOpenCommand(door *GarageDoor) *DoorOpenCommand {
	return &DoorOpenCommand{Door: door}
}

// DoorCloseCommand 关门命令
type DoorCloseCommand struct {
	Door *GarageDoor
}

func (command *DoorCloseCommand) Execute() {
	command.Door.Close()
}

func NewDoorCloseCommand(door *GarageDoor) *DoorCloseCommand {
	return &DoorCloseCommand{Door: door}
}



