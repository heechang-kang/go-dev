package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.listen("tcp", :8000) //
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil{
			log.Println(err)
		}

		io.WriteString(conn, "\nFrom TCP Server\n")
		fmt.Println(conn, "how is your day?\n")
		fmt.printf(conn, "%v", "Well I hope")

		conn.Close()
	}
}