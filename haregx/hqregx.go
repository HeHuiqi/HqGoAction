package main

import (
	"regexp"
	"fmt"
)

const text  = `
My email hehuiqi@gmail.com
My email haha@qq.com
My email hehe@abc.com.cn
`
func main() {
	simpleRegex()
}
func publicRegex()  {
	//regx := `[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`
	// ()表示正则提取
	regx := `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`

	re :=regexp.MustCompile(regx)

	//match := re.FindString(text)
	//match := re.FindAllString(text,-1)

	match := re.FindAllStringSubmatch(text,-1)
	for _,m := range match {
		fmt.Println(m)
	}
	fmt.Println(match)
}
func simpleRegex()  {

	testText := `<tr>
                           <td><span class="label">性别：</span><span field="">女</span></td>
                           <td><span class="label">生肖：</span><span field="">蛇</span></td>
                       </tr>
                       <tr>`
	regx := `<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`
	re :=regexp.MustCompile(regx)

	match := re.FindStringSubmatch(testText)
	fmt.Println(len(match))
	for _,m := range match {
		fmt.Println(m)
	}
}