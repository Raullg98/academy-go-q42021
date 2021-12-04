package util

import "testing"

func TestValidateDateFormat(t *testing.T) {

	var validateDateFormatTests = []struct {
		arg      string
		expected bool
	}{
		{"2021-12-10", true},
		{"0000-00-00", true},
		{"2020-04-02", true},
		{"200-20-20", false},
		{"2020--01", false},
		{"2000 10 01", false},
	}

	for _, test := range validateDateFormatTests {
		if output := ValidateDateFormat(test.arg); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
