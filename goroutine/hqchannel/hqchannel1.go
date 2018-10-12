package hqchannel

import (
	"sync"
	"fmt"
)

type Worker1 struct {
	in chan int
	wg *sync.WaitGroup
}

func doWork1(id int, c chan int,wg *sync.WaitGroup)  {
	for n := range c {
		fmt.Printf("Worker1 %d received %d\n",id,n)
		//结束goroutine
		wg.Done()

	}
}
func createDoWorker1(id int,wg *sync.WaitGroup) Worker1 {
	w := Worker1{in:make(chan int),
		wg:wg,}
	go doWork1(id,w.in,wg)
	return w
}

func ChannelDoWorkDemo1()  {

	count := 5

	var wg sync.WaitGroup

	var workers [5]Worker1

	for i:=0; i<count;i++{
		workers[i] = createDoWorker1(i,&wg)
	}
	wg.Add(count*2)//两次写入10个goroutine

	//使用WaitGroup
	for i,worker1 := range workers {
		worker1.in <- i+200

	}
	for i,worker1 := range workers{
		worker1.in <- i
	}

	//等待所有goroutine完成，这里会阻塞
	wg.Wait()



}