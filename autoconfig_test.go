package shims

import "testing"

func Test_DefaultValue(t *testing.T) {
	{
		v := ""
		x := DefaultValue(v, "y")
		if x != "y" {
			t.Error("Did not identify empty string value")
		}
	}
	{
		v := 0
		if DefaultValue(v, 1) != 1 {
			t.Error("Did not identify empty numeric value")
		}
	}
}
