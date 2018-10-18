package hqchannel

import (
	"fmt"
	"time"
)


type Worker struct {
	in chan int
	done chan bool
}
func doWork(id int, c chan int,done chan bool)  {
	for n := range c {
		fmt.Printf("Worker %d received %d\n",id,n)
		done <- true

	}
}
func createDoWorker(id int) Worker {
	w := Worker{in:make(chan int),
	done:make(chan bool),}
	go doWork(id,w.in,w.done)
	return w
}
//封装goroutine完成（done）通知
func ChannelDoWorkDemo()  {

	count := 5
	var workers [5]Worker

	for i:=0; i<count;i++{
		workers[i] = createDoWorker(i)
	}

	//写入值
	for i:=0; i<count;i++  {
		workers[i].in <- i
	}

	//通知写入完成
	for _,w:= range workers {
		<- w.done
	}
	//第二轮写入
	for i:=0; i<count;i++  {
		workers[i].in <- i+200

	}
	//通知写入完成
	for _,w:= range workers {
		<- w.done
	}



}
//基本用法
func ChannelDemo()  {


	/*
	c := make(chan  int)
	go worker(0,c)
	c <-1
	c <-2
	c <-3
	c <-4
	c <-5
	c <-6
	c <-7
	*/

	count := 5
	var channels [5]chan<- int

	for i:=0; i<count;i++{
		//channels[i] = make(chan  int)
		//go worker(i,channels[i])
		channels[i] = createWorker(i)
	}
	for i:=0; i<count;i++  {
		channels[i] <- i
	}

	time.Sleep(time.Second)

}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id,c)
	return c
}
func BufferedChannel()  {
	c := make(chan int,3)
	go worker(0,c)
	c <- 1
	c <- 2
	c <- 3
	time.Sleep(time.Second)
}
func ChannelClose()  {
	c := make(chan int)
	go worker(0,c)
	c <- 1
	c <- 2
	c <- 3
	//关闭通道
	close(c)
	time.Sleep(time.Second)
}

func worker(id int,c chan  int)  {

	//普通
	//for  {
	//	fmt.Printf("Worker %d received %d\n",id,<-c)
	//}


	//判断chan是否关闭1
	/*
	for {
		n,ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d\n",id,n)
	}
	*/

	//range自动检测channel是否close
	for n := range c {
		fmt.Printf("Worker %d received %d\n",id,n)

	}
}