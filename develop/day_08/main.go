package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func printPrompt() {
	fmt.Print("myshell> ")
}

func cd(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("cd: expected argument")
	}
	return os.Chdir(args[1])
}

func pwd() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}

func echo(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func killProcess(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("kill: expected PID")
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("kill: invalid PID")
	}
	return syscall.Kill(pid, syscall.SIGTERM)
}

func ps() error {
	cmd := exec.Command("ps", "-eo", "pid,ppid,comm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func executeCommand(cmd string, args []string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func handleInput(input string) error {
	args := strings.Fields(input)
	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case "cd":
		return cd(args)
	case "pwd":
		return pwd()
	case "echo":
		echo(args)
	case "kill":
		return killProcess(args)
	case "ps":
		return ps()
	case "\\quit":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		return executeCommand(args[0], args[1:])
	}
	return nil
}

func handlePipeline(input string) error {
	commands := strings.Split(input, "|")
	var prevCmd *exec.Cmd
	for i, cmdStr := range commands {
		cmdArgs := strings.Fields(strings.TrimSpace(cmdStr))
		cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
		if i == 0 {
			cmd.Stdin = os.Stdin
		} else {
			if prevCmd != nil {
				cmd.Stdin, _ = prevCmd.StdoutPipe()
			}
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			return err
		}
		prevCmd = cmd
	}
	if prevCmd != nil {
		return prevCmd.Wait()
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		if strings.Contains(input, "|") {
			if err := handlePipeline(input); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			if err := handleInput(input); err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}
