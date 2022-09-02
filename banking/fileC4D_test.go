package banking

import "testing"

func TestFileC4DVer(t *testing.T) {
	got := FileC4DVer()
	println(got)
	want := "FileC4D v0.0.0.1"
	if got != want {
		t.Errorf("FileC4DVer, want %q, got %q ", want, got)
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
