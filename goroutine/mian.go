package main

import (
	"runtime"
	"time"
	"fmt"
)

/*

启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。
在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。

*/
func goroutineUse()  {
	var a [2] int
	for i:= 0; i<2;i++ {
		go func(j int) {
			for {
				fmt.Printf("Hello goroutine %d\n",j)
				a[j] ++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(1*time.Millisecond)
	fmt.Printf("a%v",a)
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
	//hqselect.HqSelectDemo1()

	myChan := make(chan int)
	go testChan(myChan)
	myChan <- 10
	var numbers = [...]int{}
	//range会阻塞当前协程
	for _,v:= range numbers {
		fmt.Println(v)

	}

}
func testChan(c chan int)  {
	for i:=0;i<1000 ;i++  {
		if i < 10{
			fmt.Println("000000-----",i)

		}
		if i> 998 {
			fmt.Println("wait----",i)
		}

	}
	v := <- c
	fmt.Println("vvv----",v)

}

