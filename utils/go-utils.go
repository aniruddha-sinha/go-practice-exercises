package utils

import (
	"fmt"
	"os/exec"
)

func getCommandOutputOrErrors(cmd *exec.Cmd) {
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Output: ", string(out))
}

func ClearScreen() {
	cmd := exec.Command("sleep", "5")
	cmd2 := exec.Command("clear")
	getCommandOutputOrErrors(cmd)
	getCommandOutputOrErrors(cmd2)
}
