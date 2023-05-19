package cli

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"time"

	"github.com/gookit/color"
)

func getTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	var width, height int
	fmt.Sscanf(string(out), "%d %d", &height, &width)

	return width, height, nil
}

func Cli() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	resizeCmd := exec.Command("printf", "\033[8;0;0t")
	resizeCmd.Stdout = os.Stdout
	resizeCmd.Run()

	width, height, _ := getTerminalSize()

	restoreCmd := exec.Command("printf", "\033[8;"+strconv.Itoa(height)+";"+strconv.Itoa(width)+"t")
	restoreCmd.Stdout = os.Stdout
	restoreCmd.Run()
}

func Welcome() {
  currentUser, err := user.Current()

  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  text := "Hello, " + currentUser.Username + "... Happy hacking!\n"

  for _, char := range text {
    color.Cyan.Printf(string(char))
    time.Sleep(50 * time.Millisecond) 
  }
}

func Bye() {
  text := "Bye bye, see you soon!\n"
  for _, char := range text {
    color.Cyan.Printf(string(char))
    time.Sleep(50 * time.Millisecond) 
  }
}
