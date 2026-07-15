package main

import (
	"fmt"
	"os"
	"syrup-studios/suffering-server/cmd"
	"syrup-studios/suffering-server/utils"
)

func main() {
	if utils.IsElevated() {
		fmt.Println("Security Error")
		fmt.Println()
		fmt.Println("This application cannot be run with Administrator or root privileges.")
		fmt.Println("Running with elevated privileges is disabled for security reasons.")
		fmt.Println("Please restart the application using a standard user account.")
		os.Exit(1)
	}

	if !checkEULA() {
		fmt.Println("[ERROR] You need to agree to the EULA in order to start the SUFFERING Server")
		fmt.Println("[ERROR] Please open eula.txt, and change eula=false to eula=true, and restart")
		os.Exit(1)
	}

	cfg := LoadConfig()

	fmt.Printf("Starting SUFFERING Server: %s on port %d\n", cfg.Server.Name, cfg.Server.Port)
	LoadMods()

	// Everything is set up now — hand off to Cobra/the console
	cmd.Execute()
}
