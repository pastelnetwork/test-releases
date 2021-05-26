// +build windows

package configurer

import (
	"os"
	"path"
	"path/filepath"
	"syscall"
)

const (
	beforeVistaAppDir = "Application Data"
	sinceVistaAppDir  = "AppData/Roaming"
)

var defaultConfigPaths = []string{
	path.Join("$HOME", beforeVistaAppDir, "Pastel"),
	path.Join("$HOME", sinceVistaAppDir, "Pastel"),
	".",
}

// DefaultConfigPath returns the default config path for Windows OS.
func DefaultConfigPath(filename string) string {
	homeDir := os.UserHomeDir()
	appDir := beforeVistaAppDir

	v, _ := syscall.GetVersion()
	if v&0xff > 5 {
		appDir = sinceVistaAppDir
	}
	return filepath.Join(homeDir, filepath.FromSlash(appDir), "Pastel", filename)
}
