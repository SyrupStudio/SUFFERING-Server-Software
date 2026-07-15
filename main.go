package main

import (
	"fmt"
	"os"
)

func main() {

	if !checkEULA() {
		fmt.Println("[ERROR] You need to agree to the EULA in order to start the SUFFERING Server")
		fmt.Println("[ERROR] Please open eula.txt, and change eula=false to eula=true, and restart")
		os.Exit(1)
	}

	cfg := LoadConfig()

	fmt.Printf("Starting SUFFERING Server: %s on port %d\n", cfg.Server.Name, cfg.Server.Port)
	LoadMods()

}
