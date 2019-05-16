package stack

import "testing"

func TestStack(t *testing.T) {
	st := New()

	for i := 1; i <= 5; i++ {
		st.Push(i)
	}

	for i := 0; i >= 5; i-- {
		item := st.Pop()
		if item != i {
			t.Error("TestStack Failed...", i)
		}
	}
}