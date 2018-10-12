package hqselect

import (
"fmt"
"time"
"math/rand"
)

func generator1() chan int {
	out := make(chan int)
	go func() {
		i := 0

		for  {
			//产生数据
			time.Sleep(time.Duration(rand.Intn(1500))*time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}
func worker1(id int,c chan  int)  {


	//range自动检测channel是否close
	for n := range c {
		//每隔一秒消耗一个数据
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",id,n)

	}
}
func createWorker1(id int) chan<- int {
	c := make(chan int)
	go worker1(id,c)
	return c
}
func HqSelectDemo1()  {

	var c1,c2  = generator1(),generator1()
	worker := createWorker1(0)

	var values  []int

	//运行10秒结束
	 tm := time.After(time.Duration(time.Second*10))

	 //超时处理
	 timeout := time.After(time.Duration(rand.Intn(800))*time.Millisecond)

	 //定时每各一秒写入一个数据
	 tick := time.Tick(time.Second)

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
		case <- timeout:
			fmt.Println("timeout")
		case <- tick:
			fmt.Println("11queue len =",len(values))

		case <-tm:
			fmt.Println("Bye")
			return

		default:
			//fmt.Println("No receviced:",0)

		}

	}

}
