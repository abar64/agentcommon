package banking

import (
	"fmt"
	"testing"
)

func TestFileDartVer(t *testing.T) {
	fmt.Println("TestFileDartVer")
	got := FileDartVer()
	println(got)
	want := "FileDartVer v0.0.0.1"
	if got != want {
		t.Errorf("FileDartVer, want %q, got %q ", want, got)
	}

	fmt.Println("*********************************")
	fmt.Println("*** TODO: Uncomment demo test ***")
	fmt.Println("*********************************")

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

/*
func TestDart_StateMachine(t *testing.T) {

	/ *	Initial_State   int = -1
		Header_Record   int = 00
		Accepted_Record int = 01
		Pending_Record  int = 02
		Rejected_Record int = 03
		Trailer_Record  int = 99
		End_Of_Run      int = 9000
		sequence_Error  int = 9999 * /
	var got int

	t.Log("TestDart_StateMachine !!!!!!!!!!!")
	var want = Header_Record
	got = CheckState(Initial_State, Header_Record)
	if got != want {
		t.Errorf("Dart_StateMachine From Initial to Header, want %q, got %q ", want, got)
	}

	want = sequence_Error
	got = CheckState(Initial_State, Accepted_Record)
	if got != want {
		t.Errorf("Dart_StateMachine From Initial to Header, want %q, got %q ", want, got)
	}

	// Check expected states for records following Header_Record
	test := [5][2]int{{Header_Record, sequence_Error}, {Accepted_Record, Accepted_Record}, {Pending_Record, Pending_Record}, {Rejected_Record, Rejected_Record}, {Trailer_Record, Trailer_Record}}
	var currentState = Header_Record
	for _, row := range test {
		newState := row[0]
		want = row[1]
		got = CheckState(currentState, newState)
		//		t.Log(Header_Record, newState, want, got)
		if got != want {
			t.Errorf("Dart_StateMachine From Header_record(%d) to %d, want %d, got %d ", Header_Record, newState, want, got)
		}
	}
}
*/
