package hqchannel

import "fmt"

func HqStartTask()  {
	//在当前线程创建一个无缓冲的channel，
	taskChan := make(chan int)
	//开启一个goroutine执行任务并等待接收myChan的数据
	//这里会立马返回,在后台执行任务
	go doTask(taskChan)

	//向channel中写入数据,会阻塞当前线程，知道有其他线程从channel中取出数据
	taskChan <- 10
}
func doTask(c chan int)  {
	count := 1000
	for i:=0;i< count;i++  {
		if i < 5{
			fmt.Println("pre-----",i)

		}
		if i == 6 {
			fmt.Println("........")

		}
		if i> count-5 {
			fmt.Println("wait----",i)
		}

	}
	//接收c的数据
	chanValue := <- c
	fmt.Println("chanValue----",chanValue)

}