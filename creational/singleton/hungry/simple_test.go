package hungry

import "testing"

func TestSimpleHungrySingleton(t *testing.T) {
	t.Log(GetSimpleHungrySingletonInstance().Name)
}
