package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func makeRegistryKeys() {
	regPath := `dezeekeesdesktoplist`

	// Open or create registry key
	key, _, err := registry.CreateKey(registry.CLASSES_ROOT, regPath, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Error creating registry key:", err)
		os.Exit(1)
	}
	defer key.Close()

	// Set "URL Protocol" string value
	err = key.SetStringValue("URL Protocol", "")
	if err != nil {
		fmt.Println("Error setting registry value:", err)
		os.Exit(1)
	}

	// Create subkeys
	shellKey, _, err := registry.CreateKey(key, "shell", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Error creating shell subkey:", err)
		os.Exit(1)
	}
	defer shellKey.Close()

	openKey, _, err := registry.CreateKey(shellKey, "open", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Error creating open subkey:", err)
		os.Exit(1)
	}
	defer openKey.Close()

	commandKey, _, err := registry.CreateKey(openKey, "command", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Error creating command subkey:", err)
		os.Exit(1)
	}
	defer commandKey.Close()

	keyValue := fmt.Sprintf(`"%s\dezeekeesdesktoplist.exe" "%%1"`, installPath)

	// Set default value in "command" subkey
	err = commandKey.SetStringValue("", keyValue)
	if err != nil {
		fmt.Println("Error setting command registry value:", err)
		os.Exit(1)
	}

	fmt.Println("Registry keys and values created successfully.")
}

func GetInstallPath() {
	regPath := `dezeekeesdesktoplist\shell\open\command`

	// Open registry key
	key, err := registry.OpenKey(registry.CLASSES_ROOT, regPath, registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("Error opening registry key:", err)
		os.Exit(1)
	}
	defer key.Close()

	// Get default value from "command" subkey
	keyValue, _, err := key.GetStringValue("")
	if err != nil {
		fmt.Println("Error getting registry value:", err)
		os.Exit(1)
	}

	parts := strings.Split(keyValue, "\"")
	if len(parts) >= 3 {
		installPath = parts[1] // The second element is the install path
	} else {
		fmt.Println("Error parsing registry value")
		os.Exit(1)
	}
}
