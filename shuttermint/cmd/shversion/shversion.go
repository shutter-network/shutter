// Package shversion contains version information being set via linker flags when building via the
// Makefile
package shversion

import (
	"fmt"
	"runtime"
)

var version string = "(unknown)"

// Version returns shuttermint's version string.
func Version() string {
	var raceinfo string
	if raceDetectorEnabled {
		raceinfo = ", race detector enabled"
	}
	return fmt.Sprintf("%s (%s, %s-%s%s)", version, runtime.Version(), runtime.GOOS, runtime.GOARCH, raceinfo)
}
