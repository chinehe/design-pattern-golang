package memento

import "testing"

func TestMemento(t *testing.T) {
	file := NewFile()

	file.Append("你好\n")
	file.Append("姓名：Chine\n")
	file.Append("邮箱：chine@chine.com")
	file.Print()

	file.Modify("Chine", "chine")
	file.Print()

	file.Restore()
	file.Print()

	file.Restore()
	file.Print()

	file.Restore()
	file.Print()

	file.Restore()
	file.Print()
}
