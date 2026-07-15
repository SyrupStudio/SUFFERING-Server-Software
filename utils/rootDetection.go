//go:build linux || darwin

/*
Copyright © 2026 Syrup Studios>
*/

package utils

import "os"

func IsElevated() bool {
	return os.Getuid() == 0
}
