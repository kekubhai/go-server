package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}
type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

func newSimpleServer(add string) *simpleServer {
	serverUrl, err := url.Parse(add)
	handleErr(err)
	return &simpleServer{
		addr:  add,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}

}

type Loadbalancer struct {
	port           string
	rounRobinCount int
	servers        []Server
}

func NewLoadBalancer(port string, servers []Server) *Loadbalancer {
	return &Loadbalancer{
		port:           port,
		rounRobinCount: 0,
		servers:        servers,
	}
}
func handleErr(err error) {
	if err != nil {
		fmt.Printf("error : %v\n", err)
		os.Exit(1)
	}
}
