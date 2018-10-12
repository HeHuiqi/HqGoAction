package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
	"os"
)


type intGen func() int

//intGen实现Read接口进行使其可以当作文件来读
func (g intGen) Read(p []byte) (n int,err error)  {

	next := g()
	if next > 400{
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	// TODO:incorrect if p is too small!
	//将strings.NewReader(s)临时存储一下，读完p后在返回
	return  strings.NewReader(s).Read(p)
}

//斐波那契数列
// 1 1 2 3 5 8 13 21
//   a b
//		a,b
func fibonacci() intGen {

	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}


func printfileContents(reader io.Reader)  {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func writeFile(filename string)  {
	file,err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	//写入文件后要刷新一下，才能把内容从缓冲区写到硬盘中
	defer writer.Flush()
	f := fibonacci()
	for i := 0; i <20 ;i++  {
		fmt.Fprintln(writer,f())
	}
}
func safeWriteFile(filename string)  {
	file,err := os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)
	if err != nil {
		if PathError, ok := err.(*os.PathError); !ok {
			panic(err)
		}else {
			fmt.Println("ERR:",PathError.Op,PathError.Path,PathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	//写入文件后要刷新一下，才能把内容从缓冲区写到硬盘中
	defer writer.Flush()
	f := fibonacci()
	for i := 0; i <20 ;i++  {
		fmt.Fprintln(writer,f())
	}
}

func tryDefer()  {
	for i :=0; i<100 ; i++  {
		//defer语句计算后在输出，不是即时输出的
		defer fmt.Println(i)
		if i ==5 {
			return
		}
	}
}
func tryRecover()  {
	defer func() {
		r := recover()
		if err,ok := r.(error); ok {
			fmt.Println("Error occurred:",err)
		}else {
			panic(r)
		}
	}()
	//panic(errors.New("this is an error"))
	b := 0
	a := 5
	fmt.Println(a/b)
}

func main() {

	f := fibonacci()
	fmt.Println(f())
	/*
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	*/
	//printfileContents(f)
	//writeFile("fib.txt")

	//safeWriteFile("fib.txt")
	tryDefer()
	tryRecover()
}
