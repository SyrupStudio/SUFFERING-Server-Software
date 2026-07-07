package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func version() {
	// read the version file
	content, err := os.ReadFile("version.txt")
	if err != nil {
		log.Fatal("Failed to read version file :(")
	}

	// convert bytes to string and remove any new line
	version := strings.TrimSpace(string(content))

	fmt.Println("Version: %s\n", version)
}
