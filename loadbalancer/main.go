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
func (s *simpleServer) Address() string {return s.addr}
func (s *simpleServer) IsAlive()bool {return true}
func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request){
	s.proxy.ServeHTTP(rw, req)
}

func (lb *Loadbalancer) getNexavailableServer()  Server{
	server:=lb.servers[lb.rounRobinCount%len(lb.servers)]
	for server.IsAlive(){
		lb.rounRobinCount++;
		server=lb.servers[lb.rounRobinCount%len(lb.servers)]
	}
	lb.rounRobinCount++
	return server
}                          
func (lb *Loadbalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer :=lb.getNexavailableServer( )
	fmt.Print("forarding request to address %q\n", targetServer.Address())
	targetServer.Serve(rw, req)
}
func main() {
	servers := []Server{
		newSimpleServer( "https://www.facebook.com"),
		newSimpleServer("https://www.google.com"),
		newSimpleServer("https://www.github.com"),
	}
	lb := NewLoadBalancer("8000", servers)
	         handleRedirect:=func(rw http.ResponseWriter, req *http.Request){
				lb.serveProxy(rw ,req)

			 }
			 http.HandleFunc("/",handleRedirect)
			 fmt.Printf("server requests at locathost :%d'\n",lb.port)
			 http.ListenAndServe(":"+lb.port,nil)


}
