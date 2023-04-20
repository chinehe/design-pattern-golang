package impl

import (
	"design-pattern-golang/behavioral/visitor/simple"
	"fmt"
)

// ResourceFileCounter 文件统计器（具体观察者）
type ResourceFileCounter struct {

}

func (counter *ResourceFileCounter) Visit(resourceFile simple.ResourceFile) error {
	switch resourceFile.(type) {
	case *TxtFile:
		return CountTxtFile(resourceFile)
	case *CsvFile:
		return CountCsvFile(resourceFile)
	default:
		return fmt.Errorf("not found resource file type:%#v",resourceFile)
	}
}

func CountCsvFile(resourceFile simple.ResourceFile)error {
	fmt.Println("统计CSV文件大小...",resourceFile.(*CsvFile).Path)
	return nil
}

func CountTxtFile(resourceFile simple.ResourceFile)error {
	fmt.Println("统计TXT文件大小...",resourceFile.(*TxtFile).Path)
	return nil
}

