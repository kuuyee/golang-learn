package main

import (
	"container/list"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

var atExitHooks = list.New()
var atExitHookMutex sync.Mutex
var shouldOsExit = flag.Bool("exit", false, "在函数中调用os.Exit(),不执行defer")
var shouldPanic = flag.Bool("panic", false, "标准引发panic")
var shouldGoPanic = flag.Bool("go-panic", false, "在其它的go-routine中引发panic")

func RunAtExitHooks() {
	recoverErr := recover()
	if recoverErr != nil {
		fmt.Fprintf(os.Stderr, "引发Panic : %s\n", recoverErr)
	}
	for e := atExitHooks.Back(); e != nil; e = e.Prev() {
		hook := e.Value.(func())
		func(h func()) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Fprintf(os.Stderr, "钩子退出失败：%s\n", err)
				}
			}()
			h()
		}(hook)
	}
	if recoverErr != nil {
		panic(recoverErr)
	}
}

func AtExit(h func()) {
	atExitHookMutex.Lock()
	defer atExitHookMutex.Unlock()
	atExitHooks.PushBack(h)
}

func demo1() {
	fmt.Println("Called demo1().")
	AtExit(func() { fmt.Println("Registered in demo1().") })
}

func demo2exit() {
	fmt.Println("Registered in demo2().")
}

func demo2() {
	fmt.Println("Called demo2().")
	AtExit(demo2exit)
}

func demo_exit() {
	fmt.Println("Called demo_exit(), will os.Exit().\n")
	os.Exit(10)
}

func demo_panic() {
	fmt.Println("Raising panic. Let loose the dogs of war.\n")
	panic("恐慌爆发了！")
}

func demo_elsewhere_panic() {
	fmt.Println("1秒后在线程中引发恐慌")
	go func() {
		time.Sleep(time.Second)
		panic("side-shuffle")
	}()
	time.Sleep(time.Second * 2)
}

func main() {
	flag.Parse()
	if *shouldOsExit && *shouldPanic {
		fmt.Fprintf(os.Stderr, "不能同时exit和panic")
		os.Exit(1)
	}
	fmt.Println("开始.")
	defer RunAtExitHooks()

	AtExit(func() { fmt.Println("Registered in main().") })
	demo1()
	demo2()

	switch {
	case *shouldOsExit:
		demo_exit()
	case *shouldPanic:
		demo_panic()
	case *shouldGoPanic:
		demo_elsewhere_panic()
	}

	fmt.Println("main()结束行.")
}
