package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

/**********************************************************/
/* This is not finsihed, macOS is not properly supported. */
/*     The shitty code quality must be address as well    */
/**********************************************************/

func Sayonara() {
	fmt.Println("Sayonara!")

	filePath := "/usr/local/bin/gov"

	_, err := os.Stat(filePath)
	if err == nil {
		cmd := exec.Command("sudo", "rm", filePath)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("gov has been removed from your system.")
			fmt.Println("If you also want to remove Go, remove:")

			switch runtime.GOOS {
			case "windows":
				fmt.Println(`From your env path: 'USERPROFILE\go\bin;PATH'`)
			case "darwin", "linux":
				fmt.Println(`From your shell config: 'export PATH="$HOME/go/bin:$PATH"'`)
			}

			fmt.Println(`From your home directory: 'go-versions' and the 'go' alias`)
		}
	} else {
		fmt.Println("File does not exist.")
	}
}
