package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

//При подключению к несущ серверу программа должна заверщаться по timeout
func connect(addr string, timeout time.Duration) (net.Conn, error){
	end := time.After(timeout)
	d := net.Dialer{Timeout: timeout}

	for{
		select{
		case <-end:
			return nil, errors.New("timeout error")
		default:
			if conn, err := d.Dial("tcp", addr); err == nil{
				return conn, nil
			}
		}
	}
}
//TODO доделать ctr+D на линуксе
//input example:
//-host opennet.ru -port 80
//GET /
func main() {
	host := flag.String("host", "localhost", "host for connect")
	port := flag.String("port", "8080", "port for connect")
	timeout := flag.Duration("timeout", time.Second * 10, "timeout for connect")
	flag.Parse()

	//Строка для подключения
	addr := *host + ":" + *port

	//Подключаемся по tcp
	conn, err := connect(addr, *timeout)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	//Копируем stdin в сокет
	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	//Ответ из сокета в stdout
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
