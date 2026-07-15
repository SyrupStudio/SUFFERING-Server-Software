package main

/*
Copyright © 2026 Syrup Studios>
*/

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func checkEULA() bool {
	const filename = "eula.txt"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		createDefaultEULA(filename)
		return false
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.ToLower(line) == "eula=true" {
			return true
		}
	}

	return false
}

func createDefaultEULA(filename string) {
	content := `# By changing the setting below to TRUE you are indicating your agreement to our EULA.
# You can find the terms of the EULA online at: https://www.pancakse.dev (the eula coming soon)
# For copyright reasons, you must agree to this to run the software.
eula=false
`
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Critical error creating eula.txt: %v", err)
	}
}
