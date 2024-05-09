package memento

import (
	"fmt"
	"strings"
)

// File 文件（被记录者）
type File struct {
	content string
	memento *FileMementoStack
}

func NewFile() *File {
	return &File{
		memento: NewFileMementoStack(),
	}
}

func (file *File) Append(value string) {
	file.save()
	file.content += value
}

func (file *File) Modify(old, new string) {
	file.save()
	file.content = strings.ReplaceAll(file.content, old, new)
}

func (file *File) Print() {
	fmt.Println("------------")
	fmt.Println(file.content)
}

func (file *File) save() {
	file.memento.Push(file.content)
}

func (file *File) Restore() {
	file.content = file.memento.Pop()
}
