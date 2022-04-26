package scutil

import (
	"testing"
)

type paraGrading struct {
	dig     float64
	seq     byte
	grading int
}

type questionGrading struct {
	paraGrading
	ans string
}

func TestGrading(t *testing.T) {
	qs := []questionGrading{
		{
			paraGrading{3.1415926, ',', 3},
			"3.1415926",
		},
		{
			paraGrading{8844.48, ',', 3},
			"8,844.48",
		},
		{
			paraGrading{299792458, ' ', 4},
			"2 9979 2458",
		},
	}

	for _, q := range qs {
		actual := Grading(q.paraGrading.dig, q.paraGrading.seq, q.paraGrading.grading)
		if q.ans != actual {
			t.Errorf(`expect <%s>, but got <%s>`, q.ans, actual)
		}
	}
}

type paraCancelGrading struct {
	str string
	seq byte
}

type questionCancelGrading struct {
	paraCancelGrading
	ans float64
}

func TestCancelGrading(t *testing.T) {
	qs := []questionCancelGrading{
		{
			paraCancelGrading{"3.1415926", ','},
			3.1415926,
		},
		{
			paraCancelGrading{"8,844.48", ','},
			8844.48,
		},
		{
			paraCancelGrading{"2 9979 2458", ' '},
			299792458,
		},
	}

	for _, q := range qs {
		actual := CancelGrading(q.paraCancelGrading.str, q.paraCancelGrading.seq)
		if q.ans != actual {
			t.Errorf(`expect <%f>, but got <%f>`, q.ans, actual)
		}
	}
}
