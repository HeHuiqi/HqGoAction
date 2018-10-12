package hqchannel

import (
	"sync"
	"fmt"
)

type Worker2 struct {
	in chan int
	done func()
}

func doWork2(id int, w Worker2)  {
	for n := range w.in {
		fmt.Printf("Worker2 %d received %d\n",id,n)
		//结束goroutine
		w.done()

	}
}
func createDoWorker2(id int,wg *sync.WaitGroup) Worker2 {
	w := Worker2{in:make(chan int),
		done: func() {
			wg.Done()
		},}
	go doWork2(id,w)
	return w
}

func ChannelDoWorkDemo2()  {

	count := 5

	var wg sync.WaitGroup

	var workers [5]Worker2

	for i := 0; i<count;i++{
		workers[i] = createDoWorker2(i,&wg)
	}

	wg.Add(count*2)//两次写入了10个goroutine
	//使用WaitGroup
	for i,worker2 := range workers {
		worker2.in <- i+200

	}
	for i,worker2 := range workers{
		worker2.in <- i
	}

	//等待所有goroutine完成，这里会阻塞
	wg.Wait()



}