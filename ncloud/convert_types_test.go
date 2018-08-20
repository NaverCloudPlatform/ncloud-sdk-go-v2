package ncloud

import "testing"

func TestStringList(t *testing.T) {
	test := make([]string, 2)
	test[0] = "0"
	test[1] = "1"
	result := StringList(test)

	if len(test) != len(result) {
		t.Fatalf("Expected: %v, Actual: %v", test, StringListValue(result))
	}
}
