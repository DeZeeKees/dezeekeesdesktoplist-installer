package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

var installPath string

func main() {

	if !amAdmin() {
		fmt.Println("Please run this program as administrator")
		fmt.Scanln()
		return
	}

	fmt.Println("Welcome to the Dezeekees Desktop List installer")

	// split string on last occurence of \ and remove the last part
	installPath = "C:\\Program Files\\Dezeekees Desktop List"

	changeInstallPath()

	makeRegistryKeys()

	// wait for user input
	fmt.Println("Press enter to exit")
	fmt.Scanln()
}

func changeInstallPath() {
	tempInstallPath := installPath

	fmt.Println("Install path:", installPath)
	fmt.Println("Change install path? (y/n)")

	var changePath string
	fmt.Scanln(&changePath)

	if changePath == "n" {
		CallClear()

		_, err := os.Stat(installPath)

		// create the path if it does not exist
		if err != nil && os.IsNotExist(err) {
			err = os.MkdirAll(installPath, 0755)
			if err != nil {
				CallClear()
				fmt.Println("Error creating path:", err)
				changeInstallPath()
				return
			}
		}

		return
	}

	if changePath != "y" {
		// clear the console
		CallClear()

		fmt.Println("Invalid input")
		changeInstallPath()
		return
	}

	CallClear()

	// ask user for new install path
	fmt.Println("Enter new install path:")
	fmt.Scanln(&tempInstallPath)

	//check if path is an absolute path
	if tempInstallPath[1] != ':' || tempInstallPath[2] != '\\' {

		// clear the console
		CallClear()

		fmt.Println("Invalid path")
		changeInstallPath()
		return
	}

	// check if path is a valid windows path
	fileInfo, err := os.Stat(tempInstallPath)

	// if path is not valid, ask user to enter a new path
	if err != nil && !os.IsNotExist(err) {
		CallClear()
		fmt.Println("Error checking install path:", err)
		changeInstallPath()
		return
	}

	// if path does not exist, ask user if they want to create it
	if err != nil && os.IsNotExist(err) {
		CallClear()
		fmt.Println("Path does not exist")

		// ask user if they want to create the path
		fmt.Println("Create path? (y/n)")
		var createPath string
		fmt.Scanln(&createPath)

		if createPath == "n" {
			CallClear()
			changeInstallPath()
			return
		}

		if createPath != "y" {
			CallClear()
			fmt.Println("Invalid input")
			changeInstallPath()
			return
		}

		// create the path
		err = os.MkdirAll(tempInstallPath, 0755)
		if err != nil {
			CallClear()
			fmt.Println("Error creating path:", err)
			changeInstallPath()
			return
		}
	}

	// if path exists and is not a directory, ask user to enter a new path
	if fileInfo != nil && !fileInfo.IsDir() {
		CallClear()
		fmt.Println("Path is not a directory")
		changeInstallPath()
		return
	}

	installPath = tempInstallPath
	CallClear()
	changeInstallPath()
}

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
