package internal

import (
	"fmt"
)

var Date string
var CommitID string
var Tag string

var Version = fmt.Sprintf("go-mmuc %s (%s), built %s", Tag, CommitID, Date)

func PrintVersion() {
	fmt.Println(Version) //nolint:forbidigo
}
