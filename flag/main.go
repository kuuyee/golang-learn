//author kuuyee
package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	levelFlag = flag.Int("level", 0, "级别")
	bnFlag    int
)

func init() {
	flag.IntVar(&bnFlag, "bn", 3, "份数")
}

func main() {
	flag.Parse()
	count := len(os.Args)
	fmt.Printf("一共有%d个参数\n", count)
	fmt.Println("参数详情：")
	for i, v := range os.Args {
		fmt.Printf("参数%d : %s\n", i, v)
	}

	fmt.Println("参数值：")
	fmt.Println("级别：", *levelFlag)
	fmt.Println("份数：", bnFlag)
}
