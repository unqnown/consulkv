package check

import (
	"log"
	"strings"
)

// Fatal performs os.Exist(1) on non-nil error.
func Fatal(err error, args ...string) {
	if err == nil {
		return
	}
	log.Fatalf("%s: %v\n", strings.Join(args, " "), err)
}

// Error prints error message on non-nil error.
func Error(err error, args ...string) {
	if err == nil {
		return
	}
	log.Printf("ERROR: %s: %v\n", strings.Join(args, " "), err)
}
