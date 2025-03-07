package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    cmd := exec.Command("/bin/bash")

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    cmd.Env = []string{"name=shashank"}
    //следующая команда создаёт UTS, PID и IPC, NETWORK и USERNAMESPACES
    cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
				syscall.CLONE_NEWUTS |
				syscall.CLONE_NEWIPC |
				syscall.CLONE_NEWUSER |
				syscall.CLONE_NEWPID |
				syscall.CLONE_NEWNET, 
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

    if err := cmd.Run(); err != nil {
      fmt.Printf("Error running the /bin/bash command - %s\n", err)
      os.Exit(1)
    }
}
