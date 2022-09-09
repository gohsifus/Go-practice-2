package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type command struct {
	input  string
	output string
}

func (c *command) execute() (string, error) {
	args := strings.Split(c.input, " ")
	switch args[0] {
	case "cd":
		if len(args) > 2 {
			return "", errors.New("to many arguments")
		}
		err := os.Chdir(args[1])
		if err != nil {
			return "", err
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		c.output = dir
	case "echo":
		c.output = strings.Join(args[1:], " ")
	case "ps":
		cmd := exec.Command("tasklist")
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()
	case "kill":
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			return "", err
		}
		target, err := os.FindProcess(pid)
		if err != nil{
			return "", err
		}
		err = target.Kill()
		if err != nil{
			return "", err
		}
	case "exec":
		err := syscall.Exec(args[0], args, nil)
		if err != nil{
			return "", err
		}
	case "fork":
		/*id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if id == 0 {
			fmt.Println("In child:", id)
		} else {
			fmt.Println("In parent:", id)
		}*/
	case "\\quit":
		return "", errors.New("\\quit")
	default:
		fmt.Println(fmt.Sprintf("Команда %s - не поддерживается", args[0]))
	}

	return c.output, nil
}

func execCommands(commands []command) (string, error) {
	if len(commands) == 1 {
		return commands[0].execute()
	}

	//Выполняем команды одна за другой, output одной это input для другой
	currentCommand := 0
	for currentCommand <= len(commands)-1 {
		outputPrevious, err := commands[currentCommand].execute()
		if err != nil {
			return "", err
		}
		currentCommand++
		if currentCommand > len(commands)-1 {
			break
		}
		commands[currentCommand].input += " " + outputPrevious
	}

	return commands[currentCommand-1].output, nil
}

func parseCommands(inp string) ([]command, error) {
	commands := []command{}

	//Делим на пайпы
	pipes := strings.Split(inp, " | ")
	for _, v := range pipes {
		commands = append(commands, command{v, ""})
	}

	return commands, nil
}

func printWelcome() {
	fmt.Print("\n=====Shell > ")
}

func unixShellUtil() {
	printWelcome()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		commands := scanner.Text()
		parsedCommands, err := parseCommands(commands)
		if err != nil {
			panic(err)
		}
		out, err := execCommands(parsedCommands)
		if err != nil {
			if err.Error() == "\\quit" {
				break
			}
			fmt.Println(err)
		} else {
			fmt.Println(out)
		}
		printWelcome()
	}
}

func main() {
	unixShellUtil()
}
