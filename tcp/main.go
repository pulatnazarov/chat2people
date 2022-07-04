package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	con, er := net.Dial("tcp", "localhost:1234")
	if er != nil {
		panic(er)
	}
	defer con.Close()
	fmt.Println("successfully connected")
	for {
		fmt.Print("                                                                          tcp: ")
		name, er := bufio.NewReader(os.Stdin).ReadString('\n')
		if er != nil {
			panic(er)
		}
		b := make([]byte, 50)
		d := []byte(name)
		if _, er = con.Write(d); er != nil {
			panic(er)
		}

		if _, er = con.Read(b); er != nil {
			panic(er)
		}
		fmt.Println("server: ", string(b))
	}
}

func rightInput() string {
	line := ""
	a := bufio.NewScanner(os.Stdin)
	for a.Scan() {
		line = a.Text()
	}
	return line
}
