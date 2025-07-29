package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	address string
	proxy   *httputil.ReverseProxy
}

func (s *simpleServer) IsAlive() bool {
	return true
}

func (s *simpleServer) Address() string {
	return s.address
}

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

// func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) string {
// 	return s.address
// }

func NewServer(address string) *simpleServer {

	serverUrl, err := url.Parse(address)

	handleError(err)

	return &simpleServer{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type Loadbalancer struct {
	port            string
	roundRobinIndex int
	Servers         []Server
}

func NewLoadbalancer(port string, servers []Server) *Loadbalancer {
	return &Loadbalancer{
		port:            port,
		roundRobinIndex: 0,
		Servers:         servers,
	}
}

func (lb *Loadbalancer) getNextAvailableServer() (Server, error) {
	if len(lb.Servers) == 0 {
		return nil, fmt.Errorf("no servers available")
	}

	for i := 0; i < len(lb.Servers); i++ {
		server := lb.Servers[lb.roundRobinIndex]
		lb.roundRobinIndex = (lb.roundRobinIndex + 1) % len(lb.Servers)

		if (server).IsAlive() {
			return server, nil
		}
	}

	return nil, fmt.Errorf("no alive servers found")
}

func (lb *Loadbalancer) ServerProxy(rw http.ResponseWriter, req *http.Request) {

	server, err := lb.getNextAvailableServer()

	fmt.Println("Selected server:", server.Address())
	if err != nil {
		http.Error(rw, err.Error(), http.StatusServiceUnavailable)
		return
	}

	(server).Serve(rw, req)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {

	servers := []Server{
		NewServer("https://google.com"),
		NewServer("https://postman-echo.com"),
		NewServer("https://echo.free.beeceptor.com"),
	}

	lb := NewLoadbalancer(":8080", servers)

	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.ServerProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Println("Loadbalancer is running on port", lb.port)
	err := http.ListenAndServe(lb.port, nil)
	handleError(err)

}
