package parameter

// LightOnCommand 开灯命令
func LightOnCommand(light *Light) Command {
	return func() {
		light.On()
	}
}

// LightOffCommand 关灯命令
func LightOffCommand(light *Light) Command  {
	return func() {
		light.Off()
	}
}

// DoorOpenCommand 开门命令
func DoorOpenCommand(door *GarageDoor)Command {
	return func() {
		door.Open()
	}
}

// DoorCloseCommand 关门命令
func DoorCloseCommand(door *GarageDoor) Command {
	return func() {
		door.Close()
	}
}


