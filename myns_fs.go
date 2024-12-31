//providing rootfile system
package main

import (
        "fmt"
        "os"
        "os/exec"
        "path/filepath"
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

func pivotRoot(newroot string) error {
        putold := filepath.Join(newroot, "/.pivot_root")

        //привязывает к самой себе монтирование newroot - именно это выступает небольшой уловкой, которая требуется для
        //удовлетворения требования pivot_root в том, что newroot и putold не должны быть той же самой
        //файловой системой что и текущий корень
        if err := syscall.Mount(newroot, newroot, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil { return err }

        // создаём каталог putold
        if err := os.MkdirAll(putold, 0700); err != nil { return err }

        // вызываем pivot_root
        if err := syscall.PivotRoot(newroot, putold); err != nil { return err }

        // обеспечиваем установку нового текущего рабочего каталога
        root if err := os.Chdir("/"); err != nil { return err }

        //выполняем размонтирование putold, которая теперь пребывает в /.pivot_root putold = "/.pivot_root"
        if err := syscall.Unmount(putold, syscall.MNT_DETACH); err !=
        nil {
                return err
        }

        // удаляем putold
        if err := os.RemoveAll(putold); err != nil { return err }

        return nil
}

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

func child () {

cmd := exec.Command(os.Args[2], os.Args[3:]...)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
must(syscall.Sethostname([]byte("myhost")))

        if err := pivotRoot("/root/book_prep/rootfs"); err != nil
                { fmt.Printf("Error running pivot_root - %s\n",
                err) os.Exit(1)
        }
must(cmd.Run())
}

func must(err error) {
        if err != nil {
                 fmt.Printf("Error - %s\n", err)
        }
}
