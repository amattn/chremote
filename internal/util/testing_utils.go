package util

import (
	"fmt"
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, debugNum int64, expected, candidate interface{}, printArgs ...interface{}) {
	if expected == nil {
		if candidate != nil {
			extra := fmt.Sprintln(printArgs...)
			t.Errorf("%d Expected != Candidate, Candidate should be nil\n%s\nExpected (%T):\n%+v\nCandidate (%T):\n%+v", debugNum, extra, expected, expected, candidate, candidate)
			return
		}
	}

	isDeeplyEqual := reflect.DeepEqual(expected, candidate)
	if isDeeplyEqual == false {
		extra := fmt.Sprintln(printArgs...)
		t.Errorf("%d Expected != Candidate\n%s\nExpected (%T):\n%+v\nCandidate (%T):\n%+v", debugNum, extra, expected, expected, candidate, candidate)
	}
}

func AssertNotEqual(t *testing.T, debugNum int64, shouldNotBe, candidate interface{}, printArgs ...interface{}) {
	isDeeplyEqual := reflect.DeepEqual(shouldNotBe, candidate)
	if isDeeplyEqual == true {
		extra := fmt.Sprintln(printArgs...)
		t.Errorf("%d shouldNotBe == candidate but we want !=\n%s\nshouldNotBe (%T):\n%+v\ncandidate (%T):\n%+v", debugNum, extra, shouldNotBe, shouldNotBe, candidate, candidate)
	}
}

func AssertNoError(t *testing.T, debugNum int64, candidateErr error, printArgs ...interface{}) {
	if candidateErr == nil {
		return
	}

	rv := reflect.ValueOf(candidateErr)
	if rv.IsNil() == false {
		extra := fmt.Sprintln(printArgs...)
		t.Errorf("%d candidateErr != nil, candidateErr should be nil\n%s\nExpected:\n%+v\nCandidate (%T):\n%+v", debugNum, extra, nil, candidateErr, candidateErr)
		return
	}

}

func AssertIsNil(t *testing.T, debugNum int64, candidate interface{}, printArgs ...interface{}) {
	if candidate == nil {
		return
	}

	rv := reflect.ValueOf(candidate)
	if rv.IsNil() == false {
		extra := fmt.Sprintln(printArgs...)
		t.Errorf("%d Candidate != nil, Candidate should be nil\n%s\nExpected (%T):\n%+v\nCandidate (%T):\n%+v", debugNum, extra, nil, nil, candidate, candidate)
		return
	}
}

func AssertIsNotNil(t *testing.T, debugNum int64, candidate interface{}, printArgs ...interface{}) {
	if candidate == nil {
		extra := fmt.Sprintln(printArgs...)
		t.Errorf("%d Candidate == nil, Candidate should NOT be nil\n%s\nExpected: not nil\nCandidate (%T):\n%+v", debugNum, extra, candidate, candidate)
		return
	}

	rv := reflect.ValueOf(candidate)
	if rv.IsNil() == true {
		extra := fmt.Sprintln(printArgs...)
		t.Errorf("%d Candidate.IsNil() == true, Candidate should NOT be nil\n%s\nExpected: not nil\nCandidate (%T):\n%+v", debugNum, extra, candidate, candidate)
		return
	}

}
