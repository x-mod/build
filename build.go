package build

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

//String version
func String() string {
	return fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)
}

