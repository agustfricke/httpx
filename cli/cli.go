package cli

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"time"

	"github.com/gookit/color"
)

func Cli() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Welcome() {
  currentUser, err := user.Current()

  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  text := "Hello " + currentUser.Username + "!\n"

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
