package object

// Command 命令接口
type Command interface {
	Execute()
}