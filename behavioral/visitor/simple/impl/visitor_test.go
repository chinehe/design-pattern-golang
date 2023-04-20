package impl

import (
	"testing"
)

func TestVisitor(t *testing.T) {
	txtFile := new(TxtFile)
	txtFile.Path = "D://xxx.txt"

	csvFile := new(CsvFile)
	csvFile.Path = "D://yyy.csv"

	txtFile.Accept(new(FilePrinter))
	txtFile.Accept(new(ResourceFileCounter))

	csvFile.Accept(new(FilePrinter))
	csvFile.Accept(new(ResourceFileCounter))
}
