package impl

import (
	"design-pattern-golang/behavioral/visitor"
	"fmt"
)

// FilePrinter 文件打印器（具体观察者）
type FilePrinter struct {
}

func (printer *FilePrinter) Visit(resourceFile visitor.ResourceFile) error {
	switch resourceFile.(type) {
	case *TxtFile:
		return PrintTxtFile(resourceFile)
	case *CsvFile:
		return PrintCsvFile(resourceFile)
	default:
		return fmt.Errorf("not found resource file type:%#v", resourceFile)
	}
}

func PrintCsvFile(resourceFile visitor.ResourceFile) error {
	fmt.Println("打印CSV文件内容到控制台...", resourceFile.(*CsvFile).Path)
	return nil
}

func PrintTxtFile(resourceFile visitor.ResourceFile) error {
	fmt.Println("打印TXT文件内容到控制台...", resourceFile.(*TxtFile).Path)
	return nil
}
