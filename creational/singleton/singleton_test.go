package singleton

import (
	"fmt"
	"testing"

	"design-pattern-golang/creational/singleton/closure"
	"design-pattern-golang/creational/singleton/declareAndInit"
	"design-pattern-golang/creational/singleton/initMethod"
	"design-pattern-golang/creational/singleton/syncOnce"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%p\n", closure.GetInstance())
		fmt.Printf("%p\n", declareAndInit.GetInstance())
		fmt.Printf("%p\n", initMethod.GetInstance())
		fmt.Printf("%p\n", syncOnce.GetInstance())
	}
}
