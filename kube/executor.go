package kube

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/c-bata/kube-prompt/internal/debug"
)
func ProcessCmd(s string) *exec.Cmd {
	cli := os.Getenv("KUBE_CLI") // oc or kubectl
	if len(cli) == 0 {
		cli ="kubectl"
	}

	splitArgs := strings.Split(s, " ")
	cmd := exec.Command(cli, splitArgs[0:]...)
	return cmd
}

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}
	cmd := ProcessCmd(s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}

func ExecuteAndGetResult(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		debug.Log("you need to pass the something arguments")
		return ""
	}

	out := &bytes.Buffer{}
	cmd := ProcessCmd(s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		debug.Log(err.Error())
		return ""
	}
	r := string(out.Bytes())
	return r
}
