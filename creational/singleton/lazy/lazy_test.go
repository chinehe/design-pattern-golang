package lazy

import "testing"

func TestName(t *testing.T) {
	t.Log(GetInstance().Name)
}
