package chremotelib

import (
	"fmt"
	"time"
)

const (
	internalIdentifier           = "github.com/amattn/wdc"
	internalBuildTimestamp int64 = 1603406456
	internalBuildNumber    int64 = 16
	internalVersionString        = "0.1.2"
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
