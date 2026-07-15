package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func runShell(root *cobra.Command) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("SUFFERING Server console. Type 'help' for commands.")

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if fields[0] == "shutdown" || fields[0] == "exit" {
			fmt.Println("Shutting down...")
			return
		}

		cmdCopy := *root
		cmdCopy.SetArgs(fields)
		if err := cmdCopy.Execute(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
