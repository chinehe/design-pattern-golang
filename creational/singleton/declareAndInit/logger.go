package declareAndInit

import (
	"fmt"
	"time"
)

type MyLogger struct {
	Data string
}

func (l *MyLogger) Println(msg string) {
	fmt.Printf("%s %s\n", time.Now().Format("2006-01-02 15:04:05.999"), msg)
}
