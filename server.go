package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var ConnMap = make(map[string]net.Conn)
var UserMap = make(map[net.Conn]string)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handle(conn)
	}
	defer li.Close()
}

func handle(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	sc.Scan()
	username := sc.Text()
	fmt.Println(username + " Connected")
	ConnMap[username] = conn
	UserMap[conn] = username
	var user string
	var flag string
	sc.Scan()
	flag = sc.Text()
	fmt.Println(flag)
	for {
		if flag == "2" {
			sc.Scan()
			user = sc.Text()
			if ConnMap[user] == nil {
				fmt.Fprintln(conn, "0")
			}
		}
		println(user)
		if flag == "2" && ConnMap[user] != nil {
			fmt.Fprintln(conn, "1")
			fmt.Fprintln(ConnMap[user], UserMap[conn])
			break
		} else if flag != "2" {
			sc.Scan()
			user = sc.Text()
			if user != "" {
				break
			}
		}
	}
	scanner := bufio.NewScanner(ConnMap[user])
	for {
		go read(conn, user)
		scanner.Scan()
		text := scanner.Text()
		fmt.Fprintln(conn, user+": "+text)
	}
	defer conn.Close()
}

func read(conn net.Conn, user string) {
	scanner := bufio.NewScanner(conn)
	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "" {
			break
		}
		fmt.Fprintln(ConnMap[user], UserMap[conn]+": "+text)
	}
	conn.Close()
}
