package example

import "testing"

func TestExample(t *testing.T) {
	result := Exec()
	if result != "This is go-module-example project!" {
		t.Errorf("Failed to test example package")
	}
}
