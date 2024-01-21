package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main(){
	server()
}

func server() {
	fmt.Println("Listning on 8080")

	listener, _ := net.Listen("tcp", ":8080")

	for{
		conn, _ := listener.Accept()
		fmt.Println("Connection accepted")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		m, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println("Error reading connection")
				return
			}
		}

		_, e := conn.Write([]byte(m))
		if e != nil {
			fmt.Println("Error writing to connection")
			return
		}
		fmt.Printf("%v %q \n", conn.RemoteAddr(), m)
		conn.Write([]byte(m ))
	}
}
