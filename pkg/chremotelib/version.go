package chremotelib

import (
	"fmt"
	"time"
)

const (
	internalIdentifier           = "github.com/amattn/chremotelib"
	internalBuildTimestamp int64 = 1605038911
	internalBuildNumber    int64 = 26
	internalVersionString        = "0.2.0-a.1"
)

func BuildDate() time.Time {
	return time.Unix(internalBuildTimestamp, 0)
}

func BuildNumber() int64 {
	return internalBuildNumber
}

func Version() string {
	return internalVersionString
}

func VersionInfo() string {
	return fmt.Sprintf("%s (%v, build %v, build date:%v)", internalIdentifier, Version(), BuildNumber(), BuildDate())
}
