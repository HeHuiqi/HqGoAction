package hqrouterpro

import (
	"net/http"
	"fmt"
)

type HqHandle func(http.ResponseWriter, *http.Request)

type HQQHandler struct {
	routers map[string] HqHandle
}
func NewHQQHandler() *HQQHandler{
	return &HQQHandler{
		routers:make(map[string]HqHandle),
	}
}
func (h *HQQHandler)AddRoute(path string,handle HqHandle)  {

	h.routers[path] = handle
}
//实现了ServeHTTP接口
func (h *HQQHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	fmt.Println("path=",path)
	handle := h.routers[path]
	if handle == nil {
		http.NotFound(w,r)
	}else {
		handle(w,r)
	}
}