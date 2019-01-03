package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake

func main() {

	var st sonyflake.Settings
	st.StartTime = time.Now()
	st.MachineID = func() (u uint16, e error) {
		return 23423, nil
	}

	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}

	id, err := sf.NextID()
	if err != nil{
		panic("id not generated")
	}

	fmt.Printf("create ID : %v \n",id)
}
