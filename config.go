package main

/*
Copyright © 2026 Syrup Studios>
*/

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Name       string `yaml:"name"`
		Motd       string `yaml:"motd"`
		Port       int    `yaml:"port"`
		MaxPlayers int    `yaml:"maxPlayers"`
		Mods       bool   `yaml:"mods"`
	} `yaml:"server"`
}

func defaultConfig() Config {
	cfg := Config{}
	cfg.Server.Name = "My SUFFERING Server"
	cfg.Server.Motd = "My SUFFERING Server"
	cfg.Server.Port = 25666
	cfg.Server.MaxPlayers = 20
	cfg.Server.Mods = false
	return cfg
}

func LoadConfig() Config {
	const path = "config.yaml"

	if data, err := os.ReadFile(path); err == nil {
		var cfg Config
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			log.Fatalf("Error parsing config.yaml: %v", err)
		}
		return cfg
	} else if !os.IsNotExist(err) {
		log.Fatalf("Error reading config.yaml: %v", err)
	}

	cfg := defaultConfig()

	file, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error creating config.yaml: %v", err)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(&cfg); err != nil {
		log.Fatalf("Error encoding config.yaml: %v", err)
	}

	return cfg
}
