package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	lis, er := net.Listen("tcp", "localhost:1234")
	if er != nil {
		panic(er)
	}

	defer func() {
		er = lis.Close()
		if er != nil {
			log.Fatal()
		}
	}()

	println("waiting for connection")

	for {
		con, er := lis.Accept()
		if er != nil {
			panic(er)
		}
		fmt.Println("successfully accepted connection")
		defer con.Close()
		for {
			msg := make([]byte, 50) // ism
			if _, er = con.Read(msg); er != nil {
				panic(er)
			}
			fmt.Println("tcp: ", string(msg))

			fmt.Printf("                                                                          server: ")
			name, er := bufio.NewReader(os.Stdin).ReadString('\n')
			if er != nil {
				panic(er)
			}
			n := []byte(name)
			if _, er = con.Write(n); er != nil {
				panic(er)
			}
		}
	}

}
