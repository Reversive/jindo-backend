package router

import (
	"log"
	"net/http"
)

type Router struct {
	Handler *http.ServeMux
}

func NewRouter() *Router {
	return &Router{Handler: http.NewServeMux()}
}

func (s *Router) filterMethod(endpoint string, handler http.HandlerFunc, method string) {
	wrapper := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			log.Printf("Method not Allowed for %s", endpoint)
			return
		}
		handler(w, req)
	}
	s.Handler.HandleFunc(endpoint, wrapper)
}

func (s *Router) Get(endpoint string, handler http.HandlerFunc) {
	s.filterMethod(endpoint, handler, http.MethodGet)
}

func (s *Router) Post(endpoint string, handler http.HandlerFunc) {
	s.filterMethod(endpoint, handler, http.MethodPost)
}

func (s *Router) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.Handler))
}
