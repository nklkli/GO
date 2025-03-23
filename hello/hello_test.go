package hello

import "testing"

func TestSay(t *testing.T) {
	want := "Hello, test!"
	got := Say("test")
	if want != got {
		t.Errorf("wanted %v, got %v", want, got)
	}
}
