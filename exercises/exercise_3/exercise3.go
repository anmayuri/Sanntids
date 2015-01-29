// Go 1.2
// go run exercise3.go


package main
//package udp

import(
	"net"
	"fmt" 
	"time" 
)

var con *net.UDPConn 

func UDPListen() {
	fmt.Println("Listen UDP connection") 	
	sAddr, _:= net.ResolveUDPAddr("udp","localhost"+":20021")
	connection, _ := net.ListenUDP("udp", sAddr)
	con = connection 
	buffer := make([]byte, 1024)
	for {
		rLenght, rAddr, err := connection.ReadFromUDP(buffer)
		if err != nil{
			fmt.Println("error 1", err)
			return
		}
		fmt.Println("lenght", rLenght, "buffer", string(buffer), "adress", rAddr)
	}
}



func main(){
	go UDPListen()
	nAddr, err2 := net.ResolveUDPAddr("udp", "localhost"+":20021")
	if err2 != nil {
		fmt.Println("error 2")
		return
	}
	time.Sleep(2*time.Second)
	if err2 != nil {
		fmt.Println("con er ikke nil")
		return
	}
	for {
		con.WriteToUDP([]byte("send this\x00"), nAddr)
		time.Sleep(1*time.Second)
	}
	return 
}

