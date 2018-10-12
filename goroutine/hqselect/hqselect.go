package hqselect

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0

		for  {
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
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
func createWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %d\n",id,<-c)
		}
	}()
	return c
}
func HqSelectDemo()  {
	//fmt.Println(c1,c2)
	var c1,c2  = generator(),generator()
	worker := createWorker(0)

	var values  []int
	for  {
		// 非阻塞式接收
		var activeWorker chan <- int
		var activeValue int
		if len(values) >0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <- c1:
			//fmt.Println("Received form c1:",n)
			values = append(values,n)
		case n := <- c2:
			//fmt.Println("Received form c1:",n)
			values = append(values,n)
		case activeWorker <- activeValue:
			values = values[1:]
		default:
			//fmt.Println("No receviced:",0)

		}

	}

}
