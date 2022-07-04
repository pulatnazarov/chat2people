package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
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
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				for {
					msg := make([]byte, 50) // ism
					if _, err := con.Read(msg); err != nil {
						if errors.Is(err, io.EOF) {
							return
						}

						panic(err)
					}

					fmt.Println("tcp: ", string(msg))
				}
			}()

			go func() {
				defer wg.Done()
				for {
					fmt.Printf("server: ")
					name, er := bufio.NewReader(os.Stdin).ReadString('\n')
					if er != nil {
						panic(er)
					}
					n := []byte(name)
					if _, er = con.Write(n); er != nil {
						panic(er)
					}
				}
			}()

			wg.Wait()
		}
	}

}
