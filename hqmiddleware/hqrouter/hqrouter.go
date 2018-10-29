package hqrouter

import "net/http"

type middleware func(handler http.Handler) http.Handler

type Router struct {
	middlewareChain []func(http.Handler) http.Handler
	mux map[string] http.Handler

}

func NewRouter() *Router  {
	return &Router{
		middlewareChain:make([]func(http.Handler) http.Handler,0),
		mux:map[string]http.Handler{},

	}
}

func (r *Router) Use(m middleware)  {
	r.middlewareChain = append(r.middlewareChain,m)
}

func (r *Router) Add(router string,h http.Handler)  {
	var mergeHandler = h
	for i := len(r.middlewareChain)-1;i>=0 ;i--{
		mergeHandler = r.middlewareChain[i](mergeHandler)
	}
	r.mux[router] = mergeHandler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	handler := r.mux[path]
	if handler != nil {
		handler.ServeHTTP(w,req)
	}else {
		http.NotFound(w,req)
	}
}