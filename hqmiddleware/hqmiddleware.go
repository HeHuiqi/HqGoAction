package main

import (
	"net/http"
	"time"
	"log"
	"HqGoAction/hqmiddleware/hqrouter"
	"fmt"
)

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}
func home(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("home"))
}

//统计耗时中间件
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		log.Println("time:",timeElapsed)
	})
}
//请求日志中间件
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		log.Println("log:start-request-path:",r.URL.Path)

		// next handler
		next.ServeHTTP(wr, r)

		log.Println("log:end-request-path:",r.URL.Path)
	})
}

func main() {

	r := hqrouter.NewRouter()
	r.Use(timeMiddleware)
	r.Use(logMiddleware)
	r.Add("/",hello)
	r.Add("/home",home)

	//单个使用某个中间件
	//http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	fmt.Println("http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", r)
}

