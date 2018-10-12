package main

import (
	"net/http"
	"os"
	"HqGoAction/filelistserver/handler"
	"log"
	//http://localhost:8888/debug/pprof 查询性能报告
	_ "net/http/pprof"
)
/*
使用一下命令来查看服务的运行情况
web命令需要安装Graphviz这个图形生成库
go tool pprof http://localhost:8888/debug/pprof/profile
web
*/


type UserError interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter,r *http.Request)  {

	return func(w http.ResponseWriter, r *http.Request) {

		 defer func() {
		 	if r := recover(); r != nil {
				log.Printf("Panic:%v",r)
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		 }()

		err := handler(w,r)
		if err != nil {
			log.Printf("Error :"+ "handle request: %s",err)

			//用户错误
			if usr,ok := err.(UserError); ok {
				http.Error(w,usr.Message(),http.StatusBadRequest)
				return
			}

			//系统错误
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w,http.StatusText(code),code)

		}
	}
}
func main() {

	http.HandleFunc("/", errWrapper(handler.HqFileList))
	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}
}


