package hungry

import "testing"

func TestInitializationHungrySingleton(t *testing.T) {
	t.Log(GetInitializationHungrySingletonInstance().Name)
}
