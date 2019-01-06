package main

import (
	"fmt"
	"time"
)

var week time.Duration
func main() {
	t := time.Now()
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	//日历时间
	week = 60*60*24*7*1e9
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)

	//时间格式化
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))

	s:=t.Format("20060102")
	fmt.Println(t,"=>",s)
}
