package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {

	req,_:=http.NewRequest(http.MethodGet,"http://www.imooc.com",nil)
	req.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//resp,err :=http.DefaultClient.Do(req)
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {

			fmt.Println("Redict:",req.URL)
			return nil
		},
	}
	resp,err :=client.Do(req)
	//resp,err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s,err := httputil.DumpResponse(resp,true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

	
}
