package banking

import "testing"

func TestFileC4DHello(t *testing.T) {
	got := FileC4DHello()
	println(got)
	want := "FileC4DHello Hello, world."
	if got != want {
		t.Errorf("FileC4DHello, want %q, got %q ", want, got)
	}
}

/*
func TestC4D_Header(t *testing.T) {
	//var record = make(map[string]string)
	var record = make([]string, 0)
	got := C4D_Header(record)
	want := "FileC4DHello Hello, world."
	if got != want {
		t.Errorf("C4D_Header, want %q, got %q ", want, got)
	}
}
*/
