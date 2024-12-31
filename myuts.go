package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("/bin/bash")
	// Приводимые ниже операторы ссылаются на потоки input, output и error создаваемого процесса(cmd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//установка некой переменной среды
	cmd.Env = []string{"name=shashank"}
	// приводимая ниже команда создаёт пространство имён UTS для данного процесса
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/bash command - %s\n", err)
		os.Exit(1)
	}
}
