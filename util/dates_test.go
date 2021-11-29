package util

import (
	"testing"
	"time"
)

type GetDateFormattedTest struct {
	arg      time.Time
	expected string
}

var validateDateFormattedTests = []GetDateFormattedTest{
	{time.Date(2020, 12, 30, 0, 0, 0, 0, time.Local), "2020-12-30"},
	{time.Date(2020, 01, 01, 0, 0, 0, 0, time.Local), "2020-01-01"},
	{time.Date(2021, 06, 07, 0, 0, 0, 0, time.Local), "2021-06-07"},
	{time.Date(1998, 04, 02, 0, 0, 0, 0, time.Local), "1998-04-02"},
}

func TestGetNowDate(t *testing.T) {

	for _, test := range validateDateFormattedTests {
		if output := GetDateFormmated(test.arg); output != test.expected {
			t.Errorf("Output %s not equal to expected %s", output, test.expected)
		}
	}
}

type ValidateDateFormatTest struct {
	arg      string
	expected bool
}

var validateDateFormatTests = []ValidateDateFormatTest{
	{"2021-12-10", true},
	{"0000-00-00", true},
	{"2020-04-02", true},
	{"200-20-20", false},
	{"2020--01", false},
	{"2000 10 01", false},
}

func TestValidateDateFormat(t *testing.T) {

	for _, test := range validateDateFormatTests {
		if output := ValidateDateFormat(test.arg); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
