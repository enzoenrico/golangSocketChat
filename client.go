// Client currently sending a simple string
// Alter client to send a message with id and more info for tracking

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type messageType struct {
	id   int
	data string
}

func main() {
RECONNECT:
	for {
		fmt.Println("Connecting to server..")
		conn, err := net.Dial("tcp", ":8080")
		if err != nil {
			fmt.Println("Connection refused, trying again")
			time.Sleep(1 * time.Second)
			continue
		}

		fmt.Println("Connection accepted")
		for {
			var m string

			fmt.Print("[🫵] >")

			message := bufio.NewReader(os.Stdin)

			m, err := message.ReadString('\n')

			if err != nil {
				fmt.Println(err)
				continue RECONNECT
			}

			// msg := messageType{id: myId, data: m}

			// jsonData, err := json.Marshal(msg)

			if err != nil {
				fmt.Println(err)
				continue RECONNECT
			}

			fmt.Printf("Sending %q\n", m)
			// _, e := conn.Write([]byte(jsonData))
			// fmt.Println(msg)
			_, e := conn.Write([]byte(m))
			if e != nil {
				fmt.Println(err)
				continue RECONNECT
			}
			reader := bufio.NewReader(conn)
			m, er := reader.ReadString('\n')
			if er != nil {
				fmt.Println(err)
				continue RECONNECT
			}

			fmt.Printf("[Incoming] > %q \n", m)
		}
	}
}
