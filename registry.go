package main

import (
	"fmt"
	"os"

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
