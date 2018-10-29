package main

import (
	"net/http"
	"time"
	"log"
	"fmt"
	"HqGoAction/hqmiddleware/hqrouter"
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
		log.Println("time-start:",r.URL.Path)

		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		log.Println("time-end:",r.URL.Path)

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

	//r := hqrouterpro.NewHQQHandler()
	//r.AddRoute("/",hello)
	//r.AddRoute("/home",home)

	r := hqrouter.NewRouter()
	r.Use(logMiddleware)
	r.Use(timeMiddleware)
	r.Add("/",http.HandlerFunc(hello))
	r.Add("/home",http.HandlerFunc(home))



	//r := hqrouterpro.NewHQQHandler()
	//r.AddRoute("/",hello)
	//r.AddRoute("/home",home)

	fmt.Println("http://127.0.0.1:8080/")
	http.ListenAndServe(":8080", r)
}
