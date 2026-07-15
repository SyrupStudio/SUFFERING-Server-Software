package utils

import (
	"fmt"
	"os"
	"strings"
)

/*
Copyright © 2026 Syrup Studios>
*/

func GetVersion() (string, error) {
	content, err := os.ReadFile("version.txt")
	if err != nil {
		return "", fmt.Errorf("failed to read version file: %w", err)
	}

	version := strings.TrimSpace(string(content))
	return version, nil
}
