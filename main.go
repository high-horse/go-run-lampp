package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Change directory to /opt/lampp
	err := os.Chdir("/opt/lampp")
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Run the XAMPP manager with sudo
	cmd := exec.Command("sudo", "./manager-linux-x64.run")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("Error running XAMPP manager:", err)
		return
	}
}