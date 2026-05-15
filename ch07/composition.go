package main

import "fmt"

type Node struct {
	Name string
	Ip   string
	Port int
}

type Webserver struct {
	Node       // Embedded
	DomainName string
}

type DatabaseServer struct {
	Node   // Embedded
	Engine string
}

func main() {
	ting := Webserver{Node: Node{Name: "TING", Ip: "127.0.0.1", Port: 8080}, DomainName: "TING"}

	fmt.Println(ting)
	fmt.Println(ting.Node.Ip)
}
