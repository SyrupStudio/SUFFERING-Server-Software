package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gluau/gluau/vm"
	"gopkg.in/yaml.v3"
)

type ModManifest struct {
	ID               string `yaml:"id"`
	Name             string `yaml:"name"`
	Version          string `yaml:"version"`
	GameVersion      string `yaml:"game_version"`
	Author           string `yaml:"author"`
	Description      string `yaml:"description"`
	Environment      string `yaml:"environment"`
	ServerEntryPoint string `yaml:"server_entry_point"`
	ClientEntryPoint string `yaml:"client_entry_point"`
}

func LoadMods() {
	cfg := LoadConfig()
	if !cfg.Server.Mods {
		return
	}

	modsDir := "mods"

	if err := os.MkdirAll(modsDir, 0755); err != nil {
		log.Fatalf("Failed to create mods directory: %v", err)
	}

	entries, err := os.ReadDir(modsDir)
	if err != nil {
		log.Printf("Error reading mods directory: %v", err)
		return
	}

	fmt.Println("\n=== Loading Server-Side Mods ===")

	for _, entry := range entries {
		if entry.IsDir() {
			modFolder := filepath.Join(modsDir, entry.Name())
			manifestPath := filepath.Join(modFolder, "manifest.yaml")

			if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
				continue
			}

			manifestData, err := os.ReadFile(manifestPath)
			if err != nil {
				log.Printf("[%s] Failed to read manifest: %v", entry.Name(), err)
				continue
			}

			var manifest ModManifest
			if err := yaml.Unmarshal(manifestData, &manifest); err != nil {
				log.Printf("[%s] Malformed manifest YAML: %v", entry.Name(), err)
				continue
			}

			if manifest.Environment == "client" {
				fmt.Printf("[Mod Loader] Skipping client-only mod: %s\n", manifest.Name)
				continue
			}

			if manifest.ServerEntryPoint == "" {
				log.Printf("[%s] Missing 'server_entry_point' configuration", manifest.Name)
				continue
			}

			scriptPath := filepath.Join(modFolder, manifest.ServerEntryPoint)
			if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
				log.Printf("[%s] Entry point file '%s' not found", manifest.Name, manifest.ServerEntryPoint)
				continue
			}

			fmt.Printf("[Mod Loader] Loading: %s v%s by %s [%s environment]\n",
				manifest.Name, manifest.Version, manifest.Author, manifest.Environment)

			runLuauMod(scriptPath, manifest.Name)
		}
	}
	fmt.Println("================================\n")
}

func runLuauMod(scriptPath string, modName string) {
	L, err := vm.CreateLuaVm()
	if err != nil {
		log.Printf("[%s] Failed to create Luau VM: %v", modName, err)
		return
	}
	defer L.Close()

	scriptBytes, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Printf("[%s] Read error: %v", modName, err)
		return
	}

	luaFunc, err := L.LoadChunk(vm.ChunkOpts{
		Name: modName,
		Code: string(scriptBytes),
		// Env: envTab, // pass a custom global table here for sandboxing
	})
	if err != nil {
		log.Printf("[%s] Compilation/Syntax Error:\n%v", modName, err)
		return
	}

	if _, err := luaFunc.Call(); err != nil {
		log.Printf("[%s] Runtime Error: %v", modName, err)
		return
	}
}
