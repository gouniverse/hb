package hb

import (
	"testing"
)

// TestPod is a unit test for the NewPod and ToHTML methods of the Pod struct.
//
// Args:
//
//	t (testing.T): A testing object provided by the testing package.
//
// Returns:
//
//	None.
//
// Raises:
//
//	AssertionError: If the output of ToHTML method is not an empty string.
//
// Example:
//
//	>>> pod := NewPod()
//	>>> html := pod.ToHTML()
//	>>> assert(html == "")
func TestPod(t *testing.T) {
	pod := NewPod()
	html := pod.ToHTML()

	if html != "" {
		t.Error("Does not equal '' empty string, Output:" + html)
	}

	pod2 := NewPod().Children([]*Tag{
		NewBR(),
		NewI(),
	})
	html2 := pod2.ToHTML()

	if html2 != "<br /><i></i>" {
		t.Error("Does not equal '<br /><i></i>', Output:" + html2)
	}
}
