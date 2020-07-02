package chremotelib

import (
	"testing"
	"time"

	"github.com/amattn/chremote/internal/util"
)

func TestVersion(t *testing.T) {
	util.AssertEqual(t, 186603930, time.Unix(internalBuildTimestamp, 0), BuildDate())
	util.AssertEqual(t, 186603931, internalBuildNumber, BuildNumber())
	util.AssertEqual(t, 186603932, internalVersionString, Version())

	VersionInfo()
}

func TestCurrentVersion(t *testing.T) {
	cf := util.CurrentFunction()
	util.AssertEqual(t, 187382746, "github.com/amattn/chremote/pkg/chremotelib.TestCurrentVersion", cf)
}

func TestNothing(t *testing.T) {
	// do nothing.

	// uncomment the following line to verify test harness is working.

	// t.Error(3022615210)
}
