package main

import (
	"runtime"
	"time"
	"fmt"
	"HqGoAction/goroutine/hqselect"
)

func goroutineUse()  {
	var a [2] int
	for i:= 0; i<2;i++ {
		go func(j int) {
			for {
				//fmt.Printf("Hello goroutine %d\n",j)
				a[j] ++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}

func main() {

	//goroutineUse()
	//channelDemo()
	//bufferedChannel()

	//hqchannel.ChannelClose()
	//hqchannel.ChannelDoWorkDemo()
	//hqchannel.ChannelDoWorkDemo1()
	//hqchannel.ChannelDoWorkDemo2()
	//hqselect.HqSelectDemo()
	hqselect.HqSelectDemo1()


}

