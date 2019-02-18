package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt struct {
	value int
	lock sync.Mutex//同步锁
}

func (a *atomicInt) increment()  {

	//使用匿名函数来加锁一个代码块
	func(){
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()

}
func (a *atomicInt) get() int  {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}
func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("a==",a.get())
}
