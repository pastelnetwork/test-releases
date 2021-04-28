package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	runDir string

	// if `debugMode` is true, log record will contain `file` field, with a value that indicating where the log was called.
	debugMode bool
)

// SetDebugMode sets the state of debug mode
func SetDebugMode(isEnabled bool) {
	debugMode = isEnabled
}

// DebugMode returns the state of debug mode
func DebugMode() bool {
	return debugMode
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		file, _ = filepath.Rel(runDir, file)
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func init() {
	runDir, _ = os.Getwd()
}
