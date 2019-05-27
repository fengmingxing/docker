package main

import (
	"os/exec"
	"os"
	"syscall"
	"path"
	"fmt"
	"io/ioutil"
	"strconv"
)

const cgroupMemoryHierarchyMout = "/sys/fs/cgroup/memory"

func main(){
	fmt.Println(os.Args[0])
	if os.Args[0] == "cgroup" {
		fmt.Printf("current pid %d",syscall.Getpid())
		fmt.Println()
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{
		}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err:=cmd.Run();err !=nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cmd= exec.Command("/proc/self/exe")
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUTS|syscall.CLONE_NEWPID|syscall.CLONE_NEWNS,
		}
		cmd.Stdin = os.Stdin
                cmd.Stdout = os.Stdout
                cmd.Stderr = os.Stderr
		if err:=cmd.Run();err !=nil {
                        fmt.Println("ERROR:",err)
                        os.Exit(1)
                }else{
			fmt.Printf("%v",cmd.Process.Pid)
			os.Mkdir(path.Join(cgroupMemoryHierarchyMout,"testmemorylimit"),0755)
			ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMout,"testmemorylimit","tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
			//将容器进程加入到testmemorylimit这个cgroup下
			ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMout,"testmemorylimit","memory.limit_in_bytes"), []byte("100m"), 0644)
			
		}
		cmd.Process.Wait()

	}	




}
