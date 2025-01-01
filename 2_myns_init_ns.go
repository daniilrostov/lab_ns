package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
switch os.Args[1] {
	case "parent":
		parent()
	case "child":
		child()
	default:
		panic("help")
	}
}
// эта родительская функция осуществляется из главной программы, которая устанавливает все необходимые пространства имён
func parent() {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{"name=shashank"}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
					syscall.CLONE_NEWUTS |
					syscall.CLONE_NEWIPC |
					syscall.CLONE_NEWPID |
					syscall.CLONE_NEWNET |
					syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID: os.Getuid(),
				Size: 1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID: os.Getgid(),
				Size: 1,
			},
		},
	}
    must(cmd.Run())
}
// это дочерний процесс, который является копией своей родительской программы сам по себе.
func child () {
cmd := exec.Command(os.Args[2], os.Args[3:]...)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
//следующая команда устанавливает имя хоста. Основная идея состоит в демонстрации использования пространства имён UTS
must(syscall.Sethostname([]byte("daniil_test")))
// эта команда запускает оболочку, которая передаётся как аргумент программы
must(cmd.Run())
}

func must(err error) {
	if err != nil {
		fmt.Printf("Error - %s\n", err)
	}
}
