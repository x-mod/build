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

func init() {
	date = time.Now().Format(time.RFC822Z)
}
