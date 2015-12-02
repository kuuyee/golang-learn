package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Errorf("创建读写文件关联对象报错: %s", err)
		os.Exit(1)
	}

	defer r.Close()
	process, err := os.StartProcess("/bin/ps", []string{"-ef"}, &os.ProcAttr{Files: []*os.File{nil, w, os.Stderr}})
	if err != nil {
		fmt.Errorf("创建子进程报错: %s", err)
		os.Exit(1)
	}

	//Wait方法阻塞直到进程退出，然后返回一个描述ProcessState描述进程的
	//状态和可能的错误。Wait方法会释放绑定到进程p的所有资源。在大多数操
	//作系统中，进程p必须是当前进程的子进程，否则会返回错误。
	processState, err := process.Wait()
	if err != nil {
		fmt.Errorf("返回进程状态报错: %s", err)
		os.Exit(1)
	}

	err = process.Release()
	if err != nil {
		fmt.Errorf("子进程退出报错: %s", err)
		os.Exit(1)
	}
	fmt.Printf("子进程是否存在:%v\n", processState.Exited())
	fmt.Printf("子进程pid: %d\n", processState.Pid())
	fmt.Printf("是否成功退出 %v\n", processState.Success())

	fmt.Printf("退出时系统CPU时间: %s\n", processState.SystemTime())
	fmt.Printf("退出用户CPU时间: %s\n", processState.UserTime())

	err = process.Signal(syscall.SIGKILL)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Close()

}
