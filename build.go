package build

import (
	"fmt"
	"time"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

//String version
func String() string {
	return fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)
}

//String commit
func Commit() string {
	return commit
}

//String version
func Date() string {
	return date
}

//String version
func Version() string {
	return version
}

func init() {
	date = time.Now().Format(time.RFC822Z)
}
