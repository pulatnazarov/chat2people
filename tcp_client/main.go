package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

func main() {
	con, er := net.Dial("tcp", "localhost:1234")
	if er != nil {
		panic(er)
	}
	defer con.Close()
	fmt.Println("successfully connected")
	for {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for {
				fmt.Print("tcp: ")
				name, er := bufio.NewReader(os.Stdin).ReadString('\n')
				if er != nil {
					panic(er)
				}
				d := []byte(name)
				if _, er = con.Write(d); er != nil {
					panic(er)
				}
			}
		}()

		go func() {
			defer wg.Done()
			for {
				b := make([]byte, 50)
				if _, er = con.Read(b); er != nil {
					if errors.Is(er, io.EOF) {
						return
					}

					panic(er)
				}
				fmt.Println("server: ", string(b))
			}
		}()

		wg.Wait()
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
