package impl

import (
	"design-pattern-golang/behavioral/visitor"
)

// TxtFile TXT文件（具体被观察者）
type TxtFile struct {
	Path string
}

func (file *TxtFile) Accept(printer visitor.ResourceFileVisitor) error {
	return printer.Visit(file)
}

// CsvFile CSV文件（具体被观察者）
type CsvFile struct {
	Path string
}

func (file *CsvFile) Accept(printer visitor.ResourceFileVisitor) error {
	return printer.Visit(file)
}
